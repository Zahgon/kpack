package clusterlifecycle

import (
	"context"

	"github.com/google/go-containerregistry/pkg/authn"
	"knative.dev/pkg/controller"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	"github.com/pivotal/kpack/pkg/client/clientset/versioned"
	buildinformers "github.com/pivotal/kpack/pkg/client/informers/externalversions/build/v1alpha2"
	buildlisters "github.com/pivotal/kpack/pkg/client/listers/build/v1alpha2"
	"github.com/pivotal/kpack/pkg/reconciler"
	"github.com/pivotal/kpack/pkg/registry"
)

const (
	ReconcilerName = "Lifecycles"
	Kind           = "Lifecycle"
)

//go:generate counterfeiter . ClusterLifecycleReader
type ClusterLifecycleReader interface {
	Read(keychain authn.Keychain, clusterLifecycleSpec buildapi.ClusterLifecycleSpec) (buildapi.ResolvedClusterLifecycle, error)
}

func NewController(
	ctx context.Context,
	opt reconciler.Options,
	keychainFactory registry.KeychainFactory,
	clusterLifecycleInformer buildinformers.ClusterLifecycleInformer,
	clusterLifecycleReader ClusterLifecycleReader,
) *controller.Impl {
	_ = "STUB: not implemented"
	return nil
}

type Reconciler struct {
	Client                 versioned.Interface
	ClusterLifecycleLister buildlisters.ClusterLifecycleLister
	ClusterLifecycleReader ClusterLifecycleReader
	KeychainFactory        registry.KeychainFactory
}

func (c *Reconciler) Reconcile(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Reconciler) reconcileClusterLifecycleStatus(ctx context.Context, clusterLifecycle *buildapi.ClusterLifecycle) (*buildapi.ClusterLifecycle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Reconciler) updateClusterLifecycleStatus(ctx context.Context, desired *buildapi.ClusterLifecycle) error {
	_ = "STUB: not implemented"
	return nil
}
