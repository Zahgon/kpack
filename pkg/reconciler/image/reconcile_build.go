package image

import (
	"context"

	corev1 "k8s.io/api/core/v1"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
)

const (
	BuildRunningReason     = "BuildRunning"
	ResolverNotReadyReason = "ResolverNotReady"
	UnknownStateReason     = "UnknownState"
	BuildFailedReason      = "BuildFailed"
	UpToDateReason         = "UpToDate"
	NotUpToDateMessage     = "Builder is not up to date. The latest stack and buildpacks may not be in use."
)

func (c *Reconciler) reconcileBuild(ctx context.Context, image *buildapi.Image, latestBuild *buildapi.Build, sourceResolver *buildapi.SourceResolver, builder buildapi.BuilderResource, buildCacheName string) (buildapi.ImageStatus, error) {
	_ = "STUB: not implemented"
	return *new(buildapi.ImageStatus), nil
}

func noScheduledBuild(buildNeeded corev1.ConditionStatus, builder buildapi.BuilderResource, build *buildapi.Build, sourceResolver *buildapi.SourceResolver) corev1alpha1.Conditions {
	_ = "STUB: not implemented"
	return *new(corev1alpha1.Conditions)
}

func unknownStatusIfNil(condition *corev1alpha1.Condition) corev1.ConditionStatus {
	_ = "STUB: not implemented"
	return *new(corev1.ConditionStatus)
}

// Copies the message from the specified condition, or fills in a default message if nil.
// We should always have a message for non-successful conditions, as that conveys
// information to the user about what is expected.
func defaultMessageIfNil(condition *corev1alpha1.Condition, defaultMessage string) string {
	_ = "STUB: not implemented"
	return ""
}

func builderReadyCondition(builder buildapi.BuilderResource) corev1alpha1.Condition {
	_ = "STUB: not implemented"
	return *new(corev1alpha1.Condition)
}

func builderUpToDateCondition(builder buildapi.BuilderResource) corev1alpha1.Condition {
	_ = "STUB: not implemented"
	return *new(corev1alpha1.Condition)
}

func builderError(builder buildapi.BuilderResource) string { _ = "STUB: not implemented"; return "" }

func scheduledBuildCondition(build *buildapi.Build, builder buildapi.BuilderResource) corev1alpha1.Conditions {
	_ = "STUB: not implemented"
	return *new(corev1alpha1.Conditions)
}

func buildCounter(build *buildapi.Build) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func buildRunningCondition(build *buildapi.Build, builder buildapi.BuilderResource) corev1alpha1.Conditions {
	_ = "STUB: not implemented"
	return *new(corev1alpha1.Conditions)
}
