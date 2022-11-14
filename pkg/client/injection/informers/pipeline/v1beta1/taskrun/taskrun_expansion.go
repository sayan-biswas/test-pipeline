package taskrun

import (
	"context"
	"github.com/tektoncd/pipeline/pkg/client/cluster/informers/tekton/v1beta1"
	"github.com/tektoncd/pipeline/pkg/client/injection/informers/factory"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/injection"
	"knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterInformer(withClusterInformer)
}

// Key is used for associating the Informer inside the context.Context.
type ClusterKey struct{}

func withClusterInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := factory.Get(ctx)
	inf := f.Tekton().V1beta1().TaskRuns()
	return context.WithValue(ctx, ClusterKey{}, inf), inf.Informer()
}

// Get extracts the typed informer from the context.
func GetCluster(ctx context.Context) v1beta1.TaskRunClusterInformer {
	untyped := ctx.Value(ClusterKey{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch github.com/tektoncd/pipeline/pkg/client/informers/externalversions/pipeline/v1beta1.PipelineRunInformer from context.")
	}
	return untyped.(v1beta1.TaskRunClusterInformer)
}
