package buildchange

import buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"

func NewCommitChange(oldRevision, newRevision string) Change {
	_ = "STUB: not implemented"
	return *new(Change)
}

type commitChange struct {
	newRevision string
	oldRevision string
}

func (c commitChange) Reason() buildapi.BuildReason {
	_ = "STUB: not implemented"
	return *new(buildapi.BuildReason)
}

func (c commitChange) IsBuildRequired() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (c commitChange) Old() interface{} { _ = "STUB: not implemented"; return nil }

func (c commitChange) New() interface{} { _ = "STUB: not implemented"; return nil }

func (c commitChange) Priority() buildapi.BuildPriority {
	_ = "STUB: not implemented"
	return *new(buildapi.BuildPriority)
}
