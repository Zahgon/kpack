package build

import (
	_ "strconv"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
)

type ByCreationTimestamp []*buildapi.Build

func (o ByCreationTimestamp) Len() int      { _ = "STUB: not implemented"; return 0 }
func (o ByCreationTimestamp) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (o ByCreationTimestamp) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
