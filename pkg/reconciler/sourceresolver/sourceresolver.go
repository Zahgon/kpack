package sourceresolver

import (
	"context"

	"knative.dev/pkg/controller"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
	"github.com/pivotal/kpack/pkg/client/clientset/versioned"
	buildinformers "github.com/pivotal/kpack/pkg/client/informers/externalversions/build/v1alpha2"
	buildlisters "github.com/pivotal/kpack/pkg/client/listers/build/v1alpha2"
	"github.com/pivotal/kpack/pkg/config"
	"github.com/pivotal/kpack/pkg/reconciler"
)

const (
	ReconcilerName = "SourceResolvers"
)

//go:generate counterfeiter . Resolver
type Resolver interface {
	Resolve(context.Context, *buildapi.SourceResolver) (corev1alpha1.ResolvedSourceConfig, error)
	CanResolve(*buildapi.SourceResolver) bool
}

func NewController(
	ctx context.Context,
	opt reconciler.Options,
	sourceResolverInformer buildinformers.SourceResolverInformer,
	gitResolver Resolver,
	blobResolver Resolver,
	registryResolver Resolver,
	featureflags config.FeatureFlags,
) *controller.Impl {
	_ = "STUB: not implemented"
	return nil
}

//go:generate counterfeiter . Enqueuer
type Enqueuer interface {
	Enqueue(*buildapi.SourceResolver) error
}

type Reconciler struct {
	GitResolver          Resolver
	BlobResolver         Resolver
	RegistryResolver     Resolver
	Enqueuer             Enqueuer
	Client               versioned.Interface
	SourceResolverLister buildlisters.SourceResolverLister
	FeatureFlags         config.FeatureFlags
}

func (c *Reconciler) Reconcile(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Reconciler) sourceReconciler(sourceResolver *buildapi.SourceResolver) (Resolver, error) {
	_ = "STUB: not implemented"
	return *new(Resolver), nil
}

func (c *Reconciler) updateStatus(ctx context.Context, desired *buildapi.SourceResolver) error {
	_ = "STUB: not implemented"
	return nil
}

func (*Reconciler) retainCommitIfTreeUnchanged(original buildapi.SourceResolverStatus, desired *buildapi.SourceResolverStatus) {
	_ = "STUB: not implemented"
	return
}
