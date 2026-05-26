package buildchange

import (
	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
)

func NewStackChange(oldRunImageRefStr, newRunImageRefStr string) Change {
	_ = "STUB: not implemented"
	return *new(Change)
}

type stackChange struct {
	oldRunImageDigest string
	newRunImageDigest string
	err               error
}

func (s stackChange) Reason() buildapi.BuildReason {
	_ = "STUB: not implemented"
	return *new(buildapi.BuildReason)
}

func (s stackChange) IsBuildRequired() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (s stackChange) Old() interface{} { _ = "STUB: not implemented"; return nil }

func (s stackChange) New() interface{} { _ = "STUB: not implemented"; return nil }

func (s stackChange) Priority() buildapi.BuildPriority {
	_ = "STUB: not implemented"
	return *new(buildapi.BuildPriority)
}
