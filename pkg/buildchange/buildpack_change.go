package buildchange

import (
	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
)

func NewBuildpackChange(oldBuildpacks, newBuildpacks []corev1alpha1.BuildpackInfo) Change {
	_ = "STUB: not implemented"
	return *new(Change)
}

type buildpackChange struct {
	old []corev1alpha1.BuildpackInfo
	new []corev1alpha1.BuildpackInfo
}

func (b buildpackChange) Reason() buildapi.BuildReason {
	_ = "STUB: not implemented"
	return *new(buildapi.BuildReason)
}

func (b buildpackChange) IsBuildRequired() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (b buildpackChange) Old() interface{} { _ = "STUB: not implemented"; return nil }

func (b buildpackChange) New() interface{} { _ = "STUB: not implemented"; return nil }

func (b buildpackChange) Priority() buildapi.BuildPriority {
	_ = "STUB: not implemented"
	return *new(buildapi.BuildPriority)
}
