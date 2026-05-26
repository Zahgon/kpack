package git

import (
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"

	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
	"github.com/pivotal/kpack/pkg/config"
)

const defaultRemote = "origin"

type remoteGitResolver struct {
	featureFlags config.FeatureFlags
}

func (r *remoteGitResolver) Resolve(auth transport.AuthMethod, sourceConfig corev1alpha1.SourceConfig) (corev1alpha1.ResolvedSourceConfig, error) {
	_ = "STUB: not implemented"
	return *new(corev1alpha1.ResolvedSourceConfig), nil
}

func (r *remoteGitResolver) ResolveByListingRemote(auth transport.AuthMethod, sourceConfig corev1alpha1.SourceConfig) (corev1alpha1.ResolvedSourceConfig, error) {
	_ = "STUB: not implemented"
	return *new(corev1alpha1.ResolvedSourceConfig), nil
}

func sourceType(reference *plumbing.Reference) corev1alpha1.GitSourceKind {
	_ = "STUB: not implemented"
	return *new(corev1alpha1.GitSourceKind)
}
