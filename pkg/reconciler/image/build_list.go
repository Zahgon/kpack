package image

import (
	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
)

type buildList struct {
	successfulBuilds []*buildapi.Build
	failedBuilds     []*buildapi.Build
	lastBuild        *buildapi.Build
}

func newBuildList(builds []*buildapi.Build) (buildList, error) {
	_ = "STUB: not implemented"
	return *new(buildList), nil
}

//nobody enforcing this

func (l buildList) NumberFailedBuilds() int64 { _ = "STUB: not implemented"; return 0 }

func (l buildList) OldestFailure() *buildapi.Build { _ = "STUB: not implemented"; return nil }

func (l buildList) NumberSuccessfulBuilds() int64 { _ = "STUB: not implemented"; return 0 }

func (l buildList) OldestSuccess() *buildapi.Build { _ = "STUB: not implemented"; return nil }
