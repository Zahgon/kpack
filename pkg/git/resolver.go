package git

import (
	"context"

	k8sclient "k8s.io/client-go/kubernetes"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
	"github.com/pivotal/kpack/pkg/config"
)

type Resolver struct {
	remoteGitResolver remoteGitResolver
	gitKeychain       *k8sGitKeychain
	featureFlags      config.FeatureFlags
}

func NewResolver(k8sClient k8sclient.Interface, sshTrustUnknownHosts bool, featureFlags config.FeatureFlags) *Resolver {
	_ = "STUB: not implemented"
	return nil
}

func (r *Resolver) Resolve(ctx context.Context, sourceResolver *buildapi.SourceResolver) (corev1alpha1.ResolvedSourceConfig, error) {
	_ = "STUB: not implemented"
	return *new(corev1alpha1.ResolvedSourceConfig), nil
}

func (*Resolver) CanResolve(sourceResolver *buildapi.SourceResolver) bool {
	_ = "STUB: not implemented"
	return false
}
