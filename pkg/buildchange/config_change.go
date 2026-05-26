package buildchange

import (
	corev1 "k8s.io/api/core/v1"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
)

func NewConfigChange(oldConfig, newConfig Config) Change {
	_ = "STUB: not implemented"
	return *new(Change)
}

type configChange struct {
	old Config
	new Config
}

type Config struct {
	Env         []corev1.EnvVar             `json:"env,omitempty"`
	Resources   corev1.ResourceRequirements `json:"resources,omitempty"`
	Services    buildapi.Services           `json:"services,omitempty"`
	CNBBindings corev1alpha1.CNBBindings    `json:"cnbBindings,omitempty"`
	Source      corev1alpha1.SourceConfig   `json:"source,omitempty"`
}

func (c configChange) Reason() buildapi.BuildReason {
	_ = "STUB: not implemented"
	return *new(buildapi.BuildReason)
}

func (c configChange) IsBuildRequired() (bool, error) {
	_ = "STUB: not implemented"
	// Git revision changes are considered as COMMIT change
	// Ignore them as part of CONFIG Change
	return false, nil
}

func (c configChange) Old() interface{} { _ = "STUB: not implemented"; return nil }

func (c configChange) New() interface{} { _ = "STUB: not implemented"; return nil }

func (c configChange) Priority() buildapi.BuildPriority {
	_ = "STUB: not implemented"
	return *new(buildapi.BuildPriority)
}
