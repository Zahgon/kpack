package buildpack

import (
	"context"

	"github.com/google/go-containerregistry/pkg/authn"
	"knative.dev/pkg/controller"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
	"github.com/pivotal/kpack/pkg/client/clientset/versioned"
	buildinformers "github.com/pivotal/kpack/pkg/client/informers/externalversions/build/v1alpha2"
	buildlisters "github.com/pivotal/kpack/pkg/client/listers/build/v1alpha2"
	"github.com/pivotal/kpack/pkg/reconciler"
	"github.com/pivotal/kpack/pkg/registry"
)

const (
	ReconcilerName = "Buildpacks"
	Kind           = "Buildpack"
)

//go:generate counterfeiter . StoreReader
type StoreReader interface {
	Read(keychain authn.Keychain, storeImages []corev1alpha1.ImageSource) ([]corev1alpha1.BuildpackStatus, error)
}

func NewController(
	ctx context.Context,
	opt reconciler.Options,
	keychainFactory registry.KeychainFactory,
	buildpackInformer buildinformers.BuildpackInformer,
	storeReader StoreReader,
) *controller.Impl {
	_ = "STUB: not implemented"
	return nil
}

type Reconciler struct {
	Client          versioned.Interface
	StoreReader     StoreReader
	BuildpackLister buildlisters.BuildpackLister
	KeychainFactory registry.KeychainFactory
}

func (c *Reconciler) Reconcile(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Reconciler) updateBuildpackStatus(ctx context.Context, desired *buildapi.Buildpack) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Reconciler) reconcileBuildpackStatus(ctx context.Context, buildpack *buildapi.Buildpack) (*buildapi.Buildpack, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
