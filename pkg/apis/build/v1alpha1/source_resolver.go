package v1alpha1

import (
	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
)

const ActivePolling = "ActivePolling"

func (sr *SourceResolver) ResolvedSource(config corev1alpha1.ResolvedSourceConfig) {
	_ = "STUB: not implemented"
	return
}

func (sr *SourceResolver) PollingReady() bool { _ = "STUB: not implemented"; return false }

func (sr *SourceResolver) Ready() bool { _ = "STUB: not implemented"; return false }

func (sr SourceResolver) IsGit() bool { _ = "STUB: not implemented"; return false }

func (sr SourceResolver) IsBlob() bool { _ = "STUB: not implemented"; return false }

func (sr SourceResolver) IsRegistry() bool { _ = "STUB: not implemented"; return false }

func (st *SourceResolver) SourceConfig() corev1alpha1.SourceConfig {
	_ = "STUB: not implemented"
	return *new(corev1alpha1.SourceConfig)
}
