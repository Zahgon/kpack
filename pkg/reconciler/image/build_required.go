package image

import (
	corev1 "k8s.io/api/core/v1"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	"github.com/pivotal/kpack/pkg/buildchange"
)

type buildRequiredResult struct {
	ConditionStatus corev1.ConditionStatus
	ReasonsStr      string
	ChangesStr      string
	PriorityClass   string
}

func newBuildRequiredResult(summary buildchange.ChangeSummary) buildRequiredResult {
	_ = "STUB: not implemented"
	return *new(buildRequiredResult)
}

func isBuildRequired(img *buildapi.Image,
	lastBuild *buildapi.Build,
	srcResolver *buildapi.SourceResolver,
	builder buildapi.BuilderResource,
) (buildRequiredResult, error) {
	_ = "STUB: not implemented"
	return *new(buildRequiredResult), nil
}

func triggerChange(lastBuild *buildapi.Build) buildchange.Change {
	_ = "STUB: not implemented"
	return *new(buildchange.Change)
}

func commitChange(lastBuild *buildapi.Build, srcResolver *buildapi.SourceResolver) buildchange.Change {
	_ = "STUB: not implemented"
	// If the lastBuild was not a Git source, then it is not a COMMIT change
	return *new(buildchange.Change)
}

func configChange(img *buildapi.Image, lastBuild *buildapi.Build, srcResolver *buildapi.SourceResolver) buildchange.Change {
	_ = "STUB: not implemented"
	return *new(buildchange.Change)
}

func buildpackChange(lastBuild *buildapi.Build, builder buildapi.BuilderResource) buildchange.Change {
	_ = "STUB: not implemented"
	return *new(buildchange.Change)
}

func stackChange(lastBuild *buildapi.Build, builder buildapi.BuilderResource) buildchange.Change {
	_ = "STUB: not implemented"
	return *new(buildchange.Change)
}

func lifecycleChange(lastBuild *buildapi.Build, builder buildapi.BuilderResource) buildchange.Change {
	_ = "STUB: not implemented"
	return *new(buildchange.Change)
}
