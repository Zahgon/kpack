package v1alpha2

import (
	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
)

type BuilderRecord struct {
	Image                   string
	Stack                   corev1alpha1.BuildStack
	Lifecycle               ResolvedClusterLifecycle
	Buildpacks              corev1alpha1.BuildpackMetadataList
	Order                   []corev1alpha1.OrderEntry
	ObservedStoreGeneration int64
	ObservedStackGeneration int64
	OS                      string
	SignaturePaths          []CosignSignature
}

func (bs *BuilderStatus) BuilderRecord(record BuilderRecord) { _ = "STUB: not implemented"; return }

func (bs *BuilderStatus) ErrorCreate(err error) { _ = "STUB: not implemented"; return }
