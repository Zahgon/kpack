package buildchange

import (
	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
)

func NewLifecycleChange(oldLifecycle, newLifecycle string) Change {
	_ = "STUB: not implemented"
	return *new(Change)
}

type lifecycleChange struct {
	oldLifecycle string
	newLifecycle string
	err          error
}

func (l lifecycleChange) Reason() buildapi.BuildReason {
	_ = "STUB: not implemented"
	return *new(buildapi.BuildReason)
}

func (l lifecycleChange) IsBuildRequired() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (l lifecycleChange) Old() interface{} { _ = "STUB: not implemented"; return nil }

func (l lifecycleChange) New() interface{} { _ = "STUB: not implemented"; return nil }

func (l lifecycleChange) Priority() buildapi.BuildPriority {
	_ = "STUB: not implemented"
	return *new(buildapi.BuildPriority)
}
