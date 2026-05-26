package cnb

import (
	"context"

	"github.com/google/go-containerregistry/pkg/authn"
	v1 "github.com/google/go-containerregistry/pkg/v1"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
	"github.com/pivotal/kpack/pkg/registry"
)

type RemoteBuildpackFetcher interface {
	BuildpackResolver
	ResolveAndFetch(context.Context, buildapi.BuilderBuildpackRef) (RemoteBuildpackInfo, error)
}

type remoteBuildpackFetcher struct {
	BuildpackResolver
	keychainFactory registry.KeychainFactory
}

func NewRemoteBuildpackFetcher(
	factory registry.KeychainFactory,
	clusterStore *buildapi.ClusterStore,
	buildpacks []*buildapi.Buildpack, clusterBuildpacks []*buildapi.ClusterBuildpack,
) RemoteBuildpackFetcher {
	_ = "STUB: not implemented"
	return *new(RemoteBuildpackFetcher)
}

func (s *remoteBuildpackFetcher) ResolveAndFetch(ctx context.Context, ref buildapi.BuilderBuildpackRef) (RemoteBuildpackInfo, error) {
	_ = "STUB: not implemented"
	return *new(RemoteBuildpackInfo), nil
}

func (s *remoteBuildpackFetcher) fetch(ctx context.Context, remoteBuildpack K8sRemoteBuildpack) (RemoteBuildpackInfo, error) {
	_ = "STUB: not implemented"
	return *new(RemoteBuildpackInfo), nil
}

// TODO: ensure there are no cycles in the buildpack graph
func (s *remoteBuildpackFetcher) layersForOrder(ctx context.Context, order corev1alpha1.Order) ([]buildpackLayer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func layerForBuildpack(keychain authn.Keychain, buildpack corev1alpha1.BuildpackStatus) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}
