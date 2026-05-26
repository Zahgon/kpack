package clusterstack

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
	ReconcilerName = "Stacks"
	Kind           = "Stack"
)

//go:generate counterfeiter . ClusterStackReader
type ClusterStackReader interface {
	Read(keychain authn.Keychain, clusterStackSpec buildapi.ClusterStackSpec) (buildapi.ResolvedClusterStack, error)
}

func NewController(
	ctx context.Context,
	opt reconciler.Options,
	keychainFactory registry.KeychainFactory,
	clusterStackInformer buildinformers.ClusterStackInformer,
	clusterStackReader ClusterStackReader) *controller.Impl {
	_ = "STUB: not implemented"
	return nil
}

type Reconciler struct {
	Client             versioned.Interface
	ClusterStackLister buildlisters.ClusterStackLister
	ClusterStackReader ClusterStackReader
	KeychainFactory    registry.KeychainFactory
}

func (c *Reconciler) Reconcile(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Reconciler) reconcileClusterStackStatus(ctx context.Context, clusterStack *buildapi.ClusterStack) (*buildapi.ClusterStack, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Reconciler) updateClusterStackStatus(ctx context.Context, desired *buildapi.ClusterStack) error {
	_ = "STUB: not implemented"
	return nil
}
