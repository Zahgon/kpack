package cnb

import (
	lifecyclebuildpack "github.com/buildpacks/lifecycle/buildpack"
	"github.com/google/go-containerregistry/pkg/authn"
	ggcrv1 "github.com/google/go-containerregistry/pkg/v1"

	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
)

type BuildMetadata struct {
	BuildpackMetadata corev1alpha1.BuildpackMetadataList `json:"buildpackMetadata"`
	LatestCacheImage  string                             `json:"latestCacheImage"`
	LatestImage       string                             `json:"latestImage"`
	StackID           string                             `json:"stackID"`
	StackRunImage     string                             `json:"stackRunImage"`
	LifecycleVersion  string                             `json:"lifecycleVersion"`
}

type ImageFetcher interface {
	Fetch(keychain authn.Keychain, repoName string) (ggcrv1.Image, string, error)
}

type RemoteMetadataRetriever struct {
	ImageFetcher ImageFetcher
}

func (r *RemoteMetadataRetriever) GetBuildMetadata(builtImageRef, cacheTag string, keychain authn.Keychain) (*BuildMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if getting cache fails, use empty cache

func (r *RemoteMetadataRetriever) getBuiltImage(tag string, keychain authn.Keychain) (builtImage, error) {
	_ = "STUB: not implemented"
	return *new(builtImage), nil
}

func (r *RemoteMetadataRetriever) getCacheImage(cacheTag string, keychain authn.Keychain) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func readBuiltImage(appImage ggcrv1.Image, appImageId string) (builtImage, error) {
	_ = "STUB: not implemented"
	return *new(builtImage), nil
}

type builtImage struct {
	identifier        string
	buildpackMetadata []lifecyclebuildpack.GroupElement
	stack             builtImageStack
	lifecycle         builtImageLifecycle
}

type builtImageLifecycle struct {
	version string
}

type appLayersMetadata struct {
	RunImage RunImageAppMetadata `json:"runImage" toml:"run-image"`
	Stack    StackMetadata       `json:"stack" toml:"stack"`
}

type RunImageAppMetadata struct {
	TopLayer  string `json:"topLayer" toml:"top-layer"`
	Reference string `json:"reference" toml:"reference"`
}

func buildMetadataFromBuiltImage(image builtImage) corev1alpha1.BuildpackMetadataList {
	_ = "STUB: not implemented"
	return *new(corev1alpha1.BuildpackMetadataList)
}

func CompressBuildMetadata(metadata *BuildMetadata) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DecompressBuildMetadata(compressedMetadata string) (*BuildMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
