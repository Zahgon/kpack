package buildchange

import (
	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
)

type Change interface {
	Reason() buildapi.BuildReason
	IsBuildRequired() (bool, error)
	Old() interface{}
	New() interface{}
	Priority() buildapi.BuildPriority
}

type GenericChange struct {
	Reason   string                 `json:"reason,omitempty"`
	Old      interface{}            `json:"old,omitempty"`
	New      interface{}            `json:"new,omitempty"`
	Priority buildapi.BuildPriority `json:"-"`
}

func newGenericChange(change Change) GenericChange {
	_ = "STUB: not implemented"
	return *new(GenericChange)
}
