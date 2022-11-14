package pipelineresource

import (
	context "context"
	"github.com/tektoncd/pipeline/pkg/client/resource/cluster/informers/tekton/v1alpha1"
	logging "knative.dev/pkg/logging"

	factory "github.com/tektoncd/pipeline/pkg/client/resource/injection/informers/factory"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
)

func init() {
	injection.Default.RegisterInformer(withClusterInformer)
}

// Key is used for associating the Informer inside the context.Context.
type ClusterKey struct{}

func withClusterInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := factory.Get(ctx)
	inf := f.Tekton().V1alpha1().PipelineResources()
	return context.WithValue(ctx, Key{}, inf), inf.Informer()
}

// Get extracts the typed informer from the context.
func GetCluster(ctx context.Context) v1alpha1.PipelineResourceClusterInformer {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch github.com/tektoncd/pipeline/pkg/client/resource/informers/externalversions/resource/v1alpha1.PipelineResourceInformer from context.")
	}
	return untyped.(v1alpha1.PipelineResourceClusterInformer)
}
