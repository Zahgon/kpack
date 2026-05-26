package clusterbuilder

import (
	"context"

	corev1 "k8s.io/api/core/v1"

	"github.com/google/go-containerregistry/pkg/authn"
	"knative.dev/pkg/controller"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	"github.com/pivotal/kpack/pkg/client/clientset/versioned"
	buildinformers "github.com/pivotal/kpack/pkg/client/informers/externalversions/build/v1alpha2"
	buildlisters "github.com/pivotal/kpack/pkg/client/listers/build/v1alpha2"
	"github.com/pivotal/kpack/pkg/cnb"
	"github.com/pivotal/kpack/pkg/reconciler"
	"github.com/pivotal/kpack/pkg/registry"
)

const (
	ReconcilerName = "ClusterBuilders"
)

type BuilderCreator interface {
	CreateBuilder(
		ctx context.Context,
		builderKeychain authn.Keychain,
		stackKeychain authn.Keychain,
		lifecycleKeychain authn.Keychain,
		fetcher cnb.RemoteBuildpackFetcher,
		clusterStack *buildapi.ClusterStack,
		clusterLifecycle *buildapi.ClusterLifecycle,
		spec buildapi.BuilderSpec,
		serviceAccountSecrets []*corev1.Secret,
		resolvedBuilderRef string,
	) (buildapi.BuilderRecord, error)
}

type Fetcher interface {
	SecretsForServiceAccount(context.Context, string, string) ([]*corev1.Secret, error)
}

func NewController(
	ctx context.Context,
	opt reconciler.Options,
	clusterBuilderInformer buildinformers.ClusterBuilderInformer,
	builderCreator BuilderCreator,
	keychainFactory registry.KeychainFactory,
	clusterStoreInformer buildinformers.ClusterStoreInformer,
	clusterBuildpackInformer buildinformers.ClusterBuildpackInformer,
	clusterStackInformer buildinformers.ClusterStackInformer,
	clusterLifecycleInformer buildinformers.ClusterLifecycleInformer,
	secretFetcher Fetcher,
) *controller.Impl {
	_ = "STUB: not implemented"
	return nil
}

type Reconciler struct {
	Client                 versioned.Interface
	ClusterBuilderLister   buildlisters.ClusterBuilderLister
	BuilderCreator         BuilderCreator
	KeychainFactory        registry.KeychainFactory
	Tracker                reconciler.Tracker
	ClusterStoreLister     buildlisters.ClusterStoreLister
	ClusterBuildpackLister buildlisters.ClusterBuildpackLister
	ClusterStackLister     buildlisters.ClusterStackLister
	ClusterLifecycleLister buildlisters.ClusterLifecycleLister
	SecretFetcher          Fetcher
}

func (c *Reconciler) Reconcile(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Reconciler) reconcileBuilder(ctx context.Context, builder *buildapi.ClusterBuilder) (buildapi.BuilderRecord, error) {
	_ = "STUB: not implemented"
	return *new(buildapi.BuilderRecord), nil
}

func (c *Reconciler) updateStatus(ctx context.Context, desired *buildapi.ClusterBuilder) error {
	_ = "STUB: not implemented"
	return nil
}

func resolveBuilderRef(builder *buildapi.ClusterBuilder) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// this happens if there is no tag
