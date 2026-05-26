package buildchange

import (
	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
)

const (
	reasonsSeparator = ","
	errorSeparator   = "\n"
)

func NewChangeProcessor() *ChangeProcessor { _ = "STUB: not implemented"; return nil }

type ChangeProcessor struct {
	changes []GenericChange
	errStrs []string
}

func (c *ChangeProcessor) Process(change Change) *ChangeProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (c *ChangeProcessor) Summarize() (ChangeSummary, error) {
	_ = "STUB: not implemented"
	return *new(ChangeSummary), nil
}

func (c *ChangeProcessor) hasChanges() bool { _ = "STUB: not implemented"; return false }

func (c *ChangeProcessor) reasonsStr() string { _ = "STUB: not implemented"; return "" }

func (c *ChangeProcessor) changesStr() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (c *ChangeProcessor) priority() buildapi.BuildPriority {
	_ = "STUB: not implemented"
	return *new(buildapi.BuildPriority)
}
