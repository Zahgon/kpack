package v1alpha1

import (
	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
)

const (
	BuilderNotFound = "BuilderNotFound"
	BuilderNotReady = "BuilderNotReady"
)

func (im *Image) BuilderNotFound() corev1alpha1.Conditions {
	_ = "STUB: not implemented"
	return *new(corev1alpha1.Conditions)
}
