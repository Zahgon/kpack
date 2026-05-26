package duckbuilder

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
)

type DuckBuilder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DuckBuilderSpec        `json:"spec"`
	Status buildapi.BuilderStatus `json:"status"`
}

func (b *DuckBuilder) GetName() string { _ = "STUB: not implemented"; return "" }

func (b *DuckBuilder) GetNamespace() string { _ = "STUB: not implemented"; return "" }

func (b *DuckBuilder) GetKind() string { _ = "STUB: not implemented"; return "" }

type DuckBuilderSpec struct {
	ImagePullSecrets []v1.LocalObjectReference
}

func (b *DuckBuilder) Ready() bool { _ = "STUB: not implemented"; return false }

func (b *DuckBuilder) UpToDate() bool { _ = "STUB: not implemented"; return false }

func (b *DuckBuilder) BuildBuilderSpec() corev1alpha1.BuildBuilderSpec {
	_ = "STUB: not implemented"
	return *new(corev1alpha1.BuildBuilderSpec)
}

func (b *DuckBuilder) BuildpackMetadata() corev1alpha1.BuildpackMetadataList {
	_ = "STUB: not implemented"
	return *new(corev1alpha1.BuildpackMetadataList)
}

func (b *DuckBuilder) RunImage() string { _ = "STUB: not implemented"; return "" }

func (b *DuckBuilder) LifecycleVersion() string { _ = "STUB: not implemented"; return "" }

func (b *DuckBuilder) ConditionReadyMessage() string { _ = "STUB: not implemented"; return "" }
