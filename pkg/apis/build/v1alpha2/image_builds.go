package v1alpha2

import (
	corev1 "k8s.io/api/core/v1"

	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
)

const (
	BuildNumberLabel     = "image.kpack.io/buildNumber"
	ImageLabel           = "image.kpack.io/image"
	ImageGenerationLabel = "image.kpack.io/imageGeneration"

	BuildReasonAnnotation  = "image.kpack.io/reason"
	BuildChangesAnnotation = "image.kpack.io/buildChanges"
	BuildNeededAnnotation  = "image.kpack.io/additionalBuildNeeded"

	BuilderNameAnnotation = "image.kpack.io/builderName"
	BuilderKindAnnotation = "image.kpack.io/builderKind"

	BuildReasonConfig    = "CONFIG"
	BuildReasonCommit    = "COMMIT"
	BuildReasonBuildpack = "BUILDPACK"
	BuildReasonStack     = "STACK"
	BuildReasonLifecycle = "LIFECYCLE"
	BuildReasonTrigger   = "TRIGGER"
)

type BuildReason string

func (im *Image) Build(sourceResolver *SourceResolver, builder BuilderResource, latestBuild *Build, reasons, changes string, nextBuildNumber int64, priorityClass string) *Build {
	_ = "STUB: not implemented"
	return nil
}

func (is *ImageSpec) NeedVolumeCache() bool { _ = "STUB: not implemented"; return false }

func (is *ImageSpec) NeedRegistryCache() bool { _ = "STUB: not implemented"; return false }

func (im *Image) getBuildCacheConfig() *BuildCacheConfig { _ = "STUB: not implemented"; return nil }

func lastBuild(latestBuild *Build) *LastBuild { _ = "STUB: not implemented"; return nil }

func (im *Image) LatestForImage(build *Build) string { _ = "STUB: not implemented"; return "" }

func (im *Image) Services() Services { _ = "STUB: not implemented"; return *new(Services) }

func (im *Image) CNBBindings() corev1alpha1.CNBBindings {
	_ = "STUB: not implemented"
	return *new(corev1alpha1.CNBBindings)
}

func (im *Image) Env() []corev1.EnvVar { _ = "STUB: not implemented"; return nil }

func (im *Image) Resources() corev1.ResourceRequirements {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceRequirements)
}

func (im *Image) Tolerations() []corev1.Toleration { _ = "STUB: not implemented"; return nil }

func (im *Image) NodeSelector() map[string]string { _ = "STUB: not implemented"; return nil }

func (im *Image) Affinity() *corev1.Affinity { _ = "STUB: not implemented"; return nil }

func (im *Image) BuildTimeout() *int64 { _ = "STUB: not implemented"; return nil }

func (im *Image) RuntimeClassName() *string { _ = "STUB: not implemented"; return nil }

func (im *Image) SchedulerName() string { _ = "STUB: not implemented"; return "" }

func (im *Image) CacheName() string { _ = "STUB: not implemented"; return "" }

func (im *Image) BuildCache() *corev1.PersistentVolumeClaim { _ = "STUB: not implemented"; return nil }

func (im *Image) SourceResolverName() string { _ = "STUB: not implemented"; return "" }

func (im *Image) SourceResolver() *SourceResolver { _ = "STUB: not implemented"; return nil }

func (im *Image) generateTags(buildNumber string) []string { _ = "STUB: not implemented"; return nil }

// We assume that if the Image Name cannot be parsed the image will not be successfully built
// in this case we can just ignore any additional image names

func (im *Image) generateBuildName(buildNumber string) string { _ = "STUB: not implemented"; return "" }

func combine(map1, map2 map[string]string) map[string]string { _ = "STUB: not implemented"; return nil }

func (im *Image) disableAdditionalImageNames() bool { _ = "STUB: not implemented"; return false }

func (is *ImageSpec) creationTime() string { _ = "STUB: not implemented"; return "" }
