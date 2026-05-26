package v1alpha1

import (
	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
)

type BuilderRecord struct {
	Image                   string
	Stack                   corev1alpha1.BuildStack
	Buildpacks              corev1alpha1.BuildpackMetadataList
	Order                   []corev1alpha1.OrderEntry
	ObservedStoreGeneration int64
	ObservedStackGeneration int64
	OS                      string
}

func (bs *BuilderStatus) BuilderRecord(record BuilderRecord) { _ = "STUB: not implemented"; return }

func (cb *BuilderStatus) ErrorCreate(err error) { _ = "STUB: not implemented"; return }
