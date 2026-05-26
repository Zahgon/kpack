package git

import (
	"regexp"

	gogit "github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/transport"

	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
)

const commitSHARegex = `^[a-f0-9]{40}$`

var commitSHAValidator = regexp.MustCompile(commitSHARegex)

func (r *remoteGitResolver) ResolveByCloning(auth transport.AuthMethod, sourceConfig corev1alpha1.SourceConfig) (corev1alpha1.ResolvedSourceConfig, error) {
	_ = "STUB: not implemented"
	// git clone
	return *new(corev1alpha1.ResolvedSourceConfig), nil
}

func hashOfSubpath(subPath string, hash plumbing.Hash, repository *gogit.Repository) string {
	_ = "STUB: not implemented"
	return ""
}

type resolverFunc func(repository *gogit.Repository, revision string) (*plumbing.Reference, plumbing.Hash, corev1alpha1.GitSourceKind, error)

var resolvers = []resolverFunc{resolveBranch, resolveTag, resolveRevision, looksLikeACommit}

func resolveBranch(repository *gogit.Repository, branch string) (*plumbing.Reference, plumbing.Hash, corev1alpha1.GitSourceKind, error) {
	_ = "STUB: not implemented"
	return nil, *new(plumbing.Hash), *new(corev1alpha1.GitSourceKind), nil
}

func resolveTag(repository *gogit.Repository, tag string) (*plumbing.Reference, plumbing.Hash, corev1alpha1.GitSourceKind, error) {
	_ = "STUB: not implemented"
	return nil, *new(plumbing.Hash), *new(corev1alpha1.GitSourceKind), nil
}

func resolveRevision(repository *gogit.Repository, revision string) (*plumbing.Reference, plumbing.Hash, corev1alpha1.GitSourceKind, error) {
	_ = "STUB: not implemented"
	return nil, *new(plumbing.Hash), *new(corev1alpha1.GitSourceKind), nil
}

func looksLikeACommit(_ *gogit.Repository, revision string) (*plumbing.Reference, plumbing.Hash, corev1alpha1.GitSourceKind, error) {
	_ = "STUB: not implemented"
	return nil, *new(plumbing.Hash), *new(corev1alpha1.GitSourceKind), nil
}
