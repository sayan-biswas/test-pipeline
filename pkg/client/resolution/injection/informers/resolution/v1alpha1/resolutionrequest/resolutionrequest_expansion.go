package resolutionrequest

import (
	context "context"

	v1alpha1 "github.com/tektoncd/pipeline/pkg/client/resolution/cluster/informers/resolution/v1alpha1"
	factory "github.com/tektoncd/pipeline/pkg/client/resolution/injection/informers/factory"
	controller "knative.dev/pkg/controller"
	injection "knative.dev/pkg/injection"
	logging "knative.dev/pkg/logging"
)

func init() {
	injection.Default.RegisterInformer(withClusterInformer)
}

// Key is used for associating the Informer inside the context.Context.
type ClusterKey struct{}

func withClusterInformer(ctx context.Context) (context.Context, controller.Informer) {
	f := factory.Get(ctx)
	inf := f.Resolution().V1alpha1().ResolutionRequests()
	return context.WithValue(ctx, Key{}, inf), inf.Informer()
}

// Get extracts the typed informer from the context.
func GetCluster(ctx context.Context) v1alpha1.ResolutionRequestClusterInformer {
	untyped := ctx.Value(Key{})
	if untyped == nil {
		logging.FromContext(ctx).Panic(
			"Unable to fetch github.com/tektoncd/pipeline/pkg/client/resolution/informers/externalversions/resolution/v1alpha1.ResolutionRequestInformer from context.")
	}
	return untyped.(v1alpha1.ResolutionRequestClusterInformer)
}
