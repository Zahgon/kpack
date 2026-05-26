package v1alpha1

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

	BuildReasonConfig    = "CONFIG"
	BuildReasonCommit    = "COMMIT"
	BuildReasonBuildpack = "BUILDPACK"
	BuildReasonStack     = "STACK"
	BuildReasonTrigger   = "TRIGGER"
)

type BuildReason string

func (im *Image) LatestForImage(build *Build) string { _ = "STUB: not implemented"; return "" }

func (im *Image) Bindings() corev1alpha1.CNBBindings {
	_ = "STUB: not implemented"
	return *new(corev1alpha1.CNBBindings)
}

func (im *Image) Env() []corev1.EnvVar { _ = "STUB: not implemented"; return nil }

func (im *Image) Resources() corev1.ResourceRequirements {
	_ = "STUB: not implemented"
	return *new(corev1.ResourceRequirements)
}

func (im *Image) CacheName() string { _ = "STUB: not implemented"; return "" }

func (im *Image) NeedCache() bool { _ = "STUB: not implemented"; return false }

func (im *Image) BuildCache() *corev1.PersistentVolumeClaim { _ = "STUB: not implemented"; return nil }

func (im *Image) SourceResolverName() string { _ = "STUB: not implemented"; return "" }

func (im *Image) SourceResolver() *SourceResolver { _ = "STUB: not implemented"; return nil }

func (im *Image) generateTags(buildNumber string) []string { _ = "STUB: not implemented"; return nil }

// We assume that if the Image Name cannot be parsed the image will not be successfully built
// in this case we can just ignore any additional image names

func (im *Image) generateBuildName(buildNumber string) string { _ = "STUB: not implemented"; return "" }

func combine(map1, map2 map[string]string) map[string]string { _ = "STUB: not implemented"; return nil }

func (im *Image) disableAdditionalImageNames() bool { _ = "STUB: not implemented"; return false }
