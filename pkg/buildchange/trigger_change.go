package buildchange

import (
	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
)

func NewTriggerChange(dateStr string) Change { _ = "STUB: not implemented"; return *new(Change) }

type triggerChange struct {
	message string
}

func (t triggerChange) Reason() buildapi.BuildReason {
	_ = "STUB: not implemented"
	return *new(buildapi.BuildReason)
}

func (t triggerChange) IsBuildRequired() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (t triggerChange) Old() interface{} { _ = "STUB: not implemented"; return nil }

func (t triggerChange) New() interface{} { _ = "STUB: not implemented"; return nil }

func (t triggerChange) Priority() buildapi.BuildPriority {
	_ = "STUB: not implemented"
	return *new(buildapi.BuildPriority)
}
