package cnb

import (
	"archive/tar"
	"io"
	"time"

	v1 "github.com/google/go-containerregistry/pkg/v1"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
)

const (
	workspaceDir               = "/workspace"
	layersDir                  = "/layers"
	cnbDir                     = "/cnb"
	platformDir                = "/platform"
	platformEnvDir             = platformDir + "/env"
	buildpacksDir              = "/cnb/buildpacks"
	orderTomlPath              = "/cnb/order.toml"
	stackTomlPath              = "/cnb/stack.toml"
	relaxedMixinMinPlatformAPI = "0.7"
)

var (
	normalizedTime        = time.Date(1980, time.January, 1, 0, 0, 1, 0, time.UTC)
	supportedPlatformApis = []string{"0.3", "0.4", "0.5", "0.6", "0.7", "0.8"}
)

type builderBlder struct {
	baseImage         v1.Image
	lifecycleLayer    v1.Layer
	LifecycleMetadata LifecycleMetadata
	stackId           string
	order             []corev1alpha1.OrderEntry
	buildpackLayers   map[DescriptiveBuildpackInfo]buildpackLayer
	cnbUserId         int
	cnbGroupId        int
	kpackVersion      string
	runImage          string
	mixins            []string
	os                string
	arch              string
	archVariant       string
	additionalLabels  map[string]string
}

func newBuilderBldr(kpackVersion string) *builderBlder { _ = "STUB: not implemented"; return nil }

func (bb *builderBlder) AddStack(baseImage v1.Image, clusterStack *buildapi.ClusterStack) error {
	_ = "STUB: not implemented"
	return nil
}

func (bb *builderBlder) AddLifecycle(lifecycleLayer v1.Layer, lifecycleMetadata LifecycleMetadata) {
	_ = "STUB: not implemented"
	return
}

func (bb *builderBlder) AddGroup(buildpacks ...RemoteBuildpackRef) {
	_ = "STUB: not implemented"
	return
}

func (bb *builderBlder) AddAdditionalLabels(additionalLabels map[string]string) {
	_ = "STUB: not implemented"
	return
}

func (bb *builderBlder) WriteableImage() (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

func (bb *builderBlder) AddRunImage(runImage string) { _ = "STUB: not implemented"; return }

func (bb *builderBlder) validateBuilder(sortedBuildpacks []DescriptiveBuildpackInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func validatePlatformApis(builderSupportedApis []string) error {
	_ = "STUB: not implemented"
	return nil
}

func relaxedMixinContract(builderSupportedApis []string) bool {
	_ = "STUB: not implemented"
	return false
}

func (bb *builderBlder) buildpacks() []DescriptiveBuildpackInfo {
	_ = "STUB: not implemented"
	return nil
}

func (bb *builderBlder) stackLayer() (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

func (bb *builderBlder) orderLayer() (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

func (bb *builderBlder) singeFileLayer(file string, contents []byte) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

func (bb *builderBlder) defaultDirsLayer() (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

func (bb *builderBlder) kpackOwnedDir(path string) *tar.Header {
	_ = "STUB: not implemented"
	return nil
}

func (bb *builderBlder) rootOwnedDir(path string) *tar.Header {
	_ = "STUB: not implemented"
	return nil
}

func (bb *builderBlder) layerWriter(fileWriter io.Writer) layerWriter {
	_ = "STUB: not implemented"
	return *new(layerWriter)
}

type layerWriter interface {
	WriteHeader(hdr *tar.Header) error
	Write(b []byte) (int, error)
	Close() error
}

func deterministicSortBySize(layers map[DescriptiveBuildpackInfo]buildpackLayer) []DescriptiveBuildpackInfo {
	_ = "STUB: not implemented"
	return nil
}

func layers(layers ...[]v1.Layer) []v1.Layer { _ = "STUB: not implemented"; return nil }

func deduplicateLayers(layers []v1.Layer) []v1.Layer { _ = "STUB: not implemented"; return nil }
