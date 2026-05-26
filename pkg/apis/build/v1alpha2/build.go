package v1alpha2

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
)

func (*Build) GetGroupVersionKind() schema.GroupVersionKind {
	_ = "STUB: not implemented"
	return *new(schema.GroupVersionKind)
}

func (b *Build) Tag() string { _ = "STUB: not implemented"; return "" }

func (b *Build) ServiceAccount() string { _ = "STUB: not implemented"; return "" }

func (b *Build) BuilderSpec() corev1alpha1.BuildBuilderSpec {
	_ = "STUB: not implemented"
	return *new(corev1alpha1.BuildBuilderSpec)
}

func (b *Build) Services() Services { _ = "STUB: not implemented"; return *new(Services) }

func (b *Build) CnbBindings() corev1alpha1.CNBBindings {
	_ = "STUB: not implemented"
	return *new(corev1alpha1.CNBBindings)
}

func (b *Build) IsRunning() bool { _ = "STUB: not implemented"; return false }

func (b *Build) BuildRef() string { _ = "STUB: not implemented"; return "" }

func (b *Build) BuildReason() string { _ = "STUB: not implemented"; return "" }

func (b *Build) BuildChanges() string { _ = "STUB: not implemented"; return "" }

func (b *Build) PriorityClassName() string { _ = "STUB: not implemented"; return "" }

func (b *Build) ImageGeneration() int64 { _ = "STUB: not implemented"; return 0 }

func (b *Build) Stack() string { _ = "STUB: not implemented"; return "" }

func (b *Build) BuiltImage() string { _ = "STUB: not implemented"; return "" }

func (b *Build) CacheImage() string { _ = "STUB: not implemented"; return "" }

func (b *Build) IsSuccess() bool { _ = "STUB: not implemented"; return false }

func (b *Build) IsFailure() bool { _ = "STUB: not implemented"; return false }

func (b *Build) PodName() string { _ = "STUB: not implemented"; return "" }

func (b *Build) MetadataReady(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func (b *Build) Finished() bool { _ = "STUB: not implemented"; return false }

func (b *Build) NotaryV1Config() *corev1alpha1.NotaryV1Config {
	_ = "STUB: not implemented"
	return nil
}

func (b *Build) DefaultProcess() string { _ = "STUB: not implemented"; return "" }

var buildSteps = map[string]struct{}{
	PrepareContainerName:    {},
	AnalyzeContainerName:    {},
	DetectContainerName:     {},
	RestoreContainerName:    {},
	BuildContainerName:      {},
	ExportContainerName:     {},
	CompletionContainerName: {},
	RebaseContainerName:     {},
}

func BuildSteps() map[string]struct{} { _ = "STUB: not implemented"; return nil }

func IsBuildStep(step string) bool { _ = "STUB: not implemented"; return false }

func (b *Build) rebasable(builderStack string) bool { _ = "STUB: not implemented"; return false }

func (b *Build) builtWithStack(runImage string) bool { _ = "STUB: not implemented"; return false }

func (b *Build) additionalBuildNeeded() bool { _ = "STUB: not implemented"; return false }

func (b *Build) builderName() string { _ = "STUB: not implemented"; return "" }

func (b *Build) builderKind() string { _ = "STUB: not implemented"; return "" }

func podCompletedWithActiveDeadline(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }
