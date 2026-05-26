package cnb

import (
	"context"

	"github.com/google/go-containerregistry/pkg/authn"
	ggcrv1 "github.com/google/go-containerregistry/pkg/v1"
	corev1 "k8s.io/api/core/v1"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
	"github.com/pivotal/kpack/pkg/cosign"
)

type RegistryClient interface {
	Fetch(keychain authn.Keychain, repoName string) (ggcrv1.Image, string, error)
	Save(keychain authn.Keychain, tag string, image ggcrv1.Image) (string, error)
}

type RemoteBuilderCreator struct {
	RegistryClient RegistryClient
	KpackVersion   string
	ImageSigner    cosign.BuilderSigner
}

func (r *RemoteBuilderCreator) CreateBuilder(
	ctx context.Context,
	builderKeychain authn.Keychain,
	stackKeychain authn.Keychain,
	lifecycleKeychain authn.Keychain,
	fetcher RemoteBuildpackFetcher,
	clusterStack *buildapi.ClusterStack,
	clusterLifecycle *buildapi.ClusterLifecycle,
	spec buildapi.BuilderSpec,
	serviceAccountSecrets []*corev1.Secret,
	resolvedBuilderRef string,
) (buildapi.BuilderRecord, error) {
	_ = "STUB: not implemented"
	return *new(buildapi.BuilderRecord), nil
}

func getLifecycleLayer(clusterLifecycle *buildapi.ClusterLifecycle, lifecycleImage ggcrv1.Image, builderBlder *builderBlder) (lifecycleLayer ggcrv1.Layer, lifecycleMetadata LifecycleMetadata, err error) {
	_ = "STUB: not implemented"
	return *new(ggcrv1.Layer), *new(LifecycleMetadata), nil
}

func platformMatches(wantOS, wantArch, wantArchVariant string, gotOS, gotArch, gotArchVariant string) bool {
	_ = "STUB: not implemented"
	return false
}

func toCNBAPISet(from buildapi.APISet) APISet { _ = "STUB: not implemented"; return *new(APISet) }

func buildpackMetadata(buildpacks []DescriptiveBuildpackInfo) corev1alpha1.BuildpackMetadataList {
	_ = "STUB: not implemented"
	return *new(corev1alpha1.BuildpackMetadataList)
}
