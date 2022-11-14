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

// clusterFakeFactoryGenerator produces a file of listers for a given GroupVersion and
// type.
type clusterFakeFactoryGenerator struct {
	generator.DefaultGen
	outputPackage string
	imports       namer.ImportTracker
	filtered      bool

	factoryInjectionPkg          string
	fakeClientInjectionPkg       string
	sharedInformerFactoryPackage string
}

var _ generator.Generator = (*clusterFakeFactoryGenerator)(nil)

func (g *clusterFakeFactoryGenerator) Filter(c *generator.Context, t *types.Type) bool {
	// We generate a single factory, so return true once.
	if !g.filtered {
		g.filtered = true
		return true
	}
	return false
}

func (g *clusterFakeFactoryGenerator) Namers(c *generator.Context) namer.NameSystems {
	return namer.NameSystems{
		"raw": namer.NewRawNamer(g.outputPackage, g.imports),
	}
}

func (g *clusterFakeFactoryGenerator) Imports(c *generator.Context) (imports []string) {
	imports = append(imports, g.imports.ImportLines()...)
	return
}

func (g *clusterFakeFactoryGenerator) GenerateType(c *generator.Context, t *types.Type, w io.Writer) error {
	sw := generator.NewSnippetWriter(w, c, "{{", "}}")

	klog.V(5).Info("processing type ", t)

	m := map[string]interface{}{
		"factoryKey": c.Universe.Type(types.Name{Package: g.factoryInjectionPkg, Name: "Key"}),
		"factoryGet": c.Universe.Function(types.Name{Package: g.factoryInjectionPkg, Name: "Get"}),
		"clientGet":  c.Universe.Function(types.Name{Package: g.fakeClientInjectionPkg, Name: "Get"}),
		"informersNewSharedInformerFactoryWithOptions": c.Universe.Function(types.Name{Package: g.sharedInformerFactoryPackage, Name: "NewSharedInformerFactoryWithOptions"}),
		"informersSharedInformerOption":                c.Universe.Function(types.Name{Package: g.sharedInformerFactoryPackage, Name: "SharedInformerOption"}),
		"injectionRegisterInformerFactory": c.Universe.Function(types.Name{
			Package: "knative.dev/pkg/injection",
			Name:    "Fake.RegisterInformerFactory",
		}),
		"controllerGetResyncPeriod": c.Universe.Type(types.Name{Package: "knative.dev/pkg/controller", Name: "GetResyncPeriod"}),
		"contextContext": c.Universe.Type(types.Name{
			Package: "context",
			Name:    "Context",
		}),
	}

	sw.Do(clusterInjectionFakeInformerFactory, m)

	return sw.Error()
}

var clusterInjectionFakeInformerFactory = `
var Get = {{.factoryGet|raw}}

func init() {
	{{.injectionRegisterInformerFactory|raw}}(withInformerFactory)
}

func withInformerFactory(ctx {{.contextContext|raw}}) {{.contextContext|raw}} {
	c := {{.clientGet|raw}}(ctx)
	opts := make([]{{.informersSharedInformerOption|raw}}, 0, 1)
	return context.WithValue(ctx, {{.factoryKey|raw}}{},
		{{.informersNewSharedInformerFactoryWithOptions|raw}}(c, {{.controllerGetResyncPeriod|raw}}(ctx), opts...))
}
`
