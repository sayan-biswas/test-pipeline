/*
Copyright 2019 The Knative Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package generators

import (
	"io"

	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
	"k8s.io/klog/v2"
)

// clusterFakeClientGenerator produces a file of listers for a given GroupVersion and
// type.
type clusterFakeClientGenerator struct {
	generator.DefaultGen
	outputPackage string
	imports       namer.ImportTracker
	filtered      bool

	fakeClientPkg      string
	clientInjectionPkg string
}

var _ generator.Generator = (*clusterFakeClientGenerator)(nil)

func (g *clusterFakeClientGenerator) Filter(c *generator.Context, t *types.Type) bool {
	// We generate a single client, so return true once.
	if !g.filtered {
		g.filtered = true
		return true
	}
	return false
}

func (g *clusterFakeClientGenerator) Namers(c *generator.Context) namer.NameSystems {
	return namer.NameSystems{
		"raw": namer.NewRawNamer(g.outputPackage, g.imports),
	}
}

func (g *clusterFakeClientGenerator) Imports(c *generator.Context) (imports []string) {
	imports = append(imports, g.imports.ImportLines()...)
	return
}

func (g *clusterFakeClientGenerator) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	sw := generator.NewSnippetWriter(w, c, "{{", "}}")

	klog.V(5).Info("processing type ", t)

	m := map[string]interface{}{
		"clientKey":  c.Universe.Type(types.Name{Package: g.clientInjectionPkg, Name: "Key"}),
		"fakeClient": c.Universe.Type(types.Name{Package: g.fakeClientPkg, Name: "ClusterClientset"}),
		"injectionRegisterClient": c.Universe.Function(types.Name{
			Package: "knative.dev/pkg/injection",
			Name:    "Fake.RegisterClient",
		}),
		"injectionRegisterClientFetcher": c.Universe.Function(types.Name{
			Package: "knative.dev/pkg/injection",
			Name:    "Fake.RegisterClientFetcher",
		}),
		"loggingFromContext": c.Universe.Function(types.Name{
			Package: "knative.dev/pkg/logging",
			Name:    "FromContext",
		}),
		"contextContext": c.Universe.Type(types.Name{
			Package: "context",
			Name:    "Context",
		}),
		"restConfig":    c.Universe.Type(types.Name{Package: "k8s.io/client-go/rest", Name: "Config"}),
		"runtimeObject": c.Universe.Type(types.Name{Package: "k8s.io/apimachinery/pkg/runtime", Name: "Object"}),
	}

	sw.Do(clusterInjectionFakeClient, m)

	return sw.Error()
}

var clusterInjectionFakeClient = `
func init() {
	{{.injectionRegisterClient|raw}}(withClient)
	{{.injectionRegisterClientFetcher|raw}}(func(ctx context.Context) interface{} {
		return Get(ctx)
	})
}

func withClient(ctx {{.contextContext|raw}}, cfg *{{.restConfig|raw}}) {{.contextContext|raw}} {
	ctx, _ = With(ctx)
	return ctx
}

func With(ctx {{.contextContext|raw}}, objects ...{{.runtimeObject|raw}}) ({{.contextContext|raw}}, *{{.fakeClient|raw}}) {
	cs := fake.NewSimpleClientset(objects...)
	return context.WithValue(ctx, {{.clientKey|raw}}{}, cs), cs
}

// Get extracts the Kubernetes client from the context.
func Get(ctx {{.contextContext|raw}}) *{{.fakeClient|raw}} {
	untyped := ctx.Value({{.clientKey|raw}}{})
	if untyped == nil {
		{{.loggingFromContext|raw}}(ctx).Panic(
			"Unable to fetch {{.fakeClient}} from context.")
	}
	return untyped.(*fake.Clientset)
}
`
