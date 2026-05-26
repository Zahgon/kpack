package buildchange

import (
	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
)

type ChangeSummary struct {
	HasChanges bool
	ReasonsStr string
	ChangesStr string
	Priority   buildapi.BuildPriority
}

func NewChangeSummary(hasChanges bool, reasonsStr, changesStr string, priority buildapi.BuildPriority) (ChangeSummary, error) {
	_ = "STUB: not implemented"
	return *new(ChangeSummary), nil
}

func (c ChangeSummary) IsValid() bool { _ = "STUB: not implemented"; return false }
