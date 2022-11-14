/*
Copyright 2021 The Knative Authors

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

package client

import (
	"context"
	"github.com/tektoncd/pipeline/pkg/client/cluster/clientset"
	"k8s.io/client-go/rest"
	"knative.dev/pkg/injection"
	"knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterClient(withClusterClientFromConfig)
}

// ClusterKey is used as the key for associating information with a context.Context.
type ClusterKey struct{}

func withClusterClientFromConfig(ctx context.Context, cfg *rest.Config) context.Context {
	return context.WithValue(ctx, ClusterKey{}, clientset.NewForConfigOrDie(cfg))
}

// GetCluster extracts the kubernetes.Interface client from the context.
func GetCluster(ctx context.Context) clientset.ClusterInterface {
	untyped := ctx.Value(ClusterKey{})
	if untyped == nil {
		if injection.GetConfig(ctx) == nil {
			logging.FromContext(ctx).Panic(
				"Unable to fetch k8s.io/client-go/kubernetes.Interface from context. This context is not the application context (which is typically given to constructors via sharedmain).")
		} else {
			logging.FromContext(ctx).Panic(
				"Unable to fetch k8s.io/client-go/kubernetes.Interface from context.")
		}
	}
	return untyped.(clientset.ClusterInterface)
}
