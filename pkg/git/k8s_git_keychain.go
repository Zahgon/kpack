package git

import (
	"context"

	"github.com/go-git/go-git/v5/plumbing/transport"
	v1 "k8s.io/api/core/v1"
	k8sclient "k8s.io/client-go/kubernetes"

	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
	"github.com/pivotal/kpack/pkg/secret"
)

type k8sGitKeychain struct {
	secretFetcher        secret.Fetcher
	sshTrustUnknownHosts bool
}

var anonymousAuth transport.AuthMethod = nil

func newK8sGitKeychain(k8sClient k8sclient.Interface, sshTrustUnknownHosts bool) *k8sGitKeychain {
	_ = "STUB: not implemented"
	return nil
}

func (k *k8sGitKeychain) Resolve(ctx context.Context, namespace, serviceAccount string, git corev1alpha1.Git) (transport.AuthMethod, error) {
	_ = "STUB: not implemented"
	return *new(transport.AuthMethod), nil
}

func fetchBasicAuth(s *v1.Secret) func() (secret.BasicAuth, error) {
	_ = "STUB: not implemented"
	return nil
}

func fetchSshAuth(s *v1.Secret) func() (secret.SSH, error) { _ = "STUB: not implemented"; return nil }

var matchingDomains = []string{
	// Allow naked domains
	"%s",
	// Allow scheme-prefixed.
	"https://%s",
	"http://%s",
	"git@%s",
}

func gitUrlMatch(urlMatch, annotatedUrl string) bool { _ = "STUB: not implemented"; return false }
