package buildpod

import (
	"context"

	"github.com/Masterminds/semver/v3"
	"github.com/google/go-containerregistry/pkg/authn"
	ggcrv1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/pivotal/kpack/pkg/client/clientset/versioned"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/dynamic"
	k8sclient "k8s.io/client-go/kubernetes"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
	"github.com/pivotal/kpack/pkg/duckprovisionedserviceable"
	"github.com/pivotal/kpack/pkg/registry"
)

const (
	builderMetadataLabel = "io.buildpacks.builder.metadata"
	cnbUserId            = "CNB_USER_ID"
	cnbGroupId           = "CNB_GROUP_ID"
)

type ImageFetcher interface {
	Fetch(keychain authn.Keychain, repoName string) (ggcrv1.Image, string, error)
}

type Generator struct {
	BuildPodConfig            buildapi.BuildPodImages
	K8sClient                 k8sclient.Interface
	KeychainFactory           registry.KeychainFactory
	ImageFetcher              ImageFetcher
	DynamicClient             dynamic.Interface
	MaximumPlatformApiVersion *semver.Version
	InjectedSidecarSupport    bool
	SSHTrustUnknownHost       bool
	KpackClient               versioned.Interface
}

type BuildPodable interface {
	GetName() string
	GetNamespace() string
	ServiceAccount() string
	BuilderSpec() corev1alpha1.BuildBuilderSpec
	CnbBindings() corev1alpha1.CNBBindings
	Services() buildapi.Services

	BuildPod(buildapi.BuildPodImages, buildapi.BuildContext) (*corev1.Pod, error)
}

func (g *Generator) Generate(ctx context.Context, build BuildPodable) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) fetchServiceBindings(ctx context.Context, build BuildPodable) ([]buildapi.ServiceBinding, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) readProvisionedServiceDuckType(ctx context.Context, build BuildPodable, s corev1.ObjectReference) (duckprovisionedserviceable.ProvisionedServicable, error) {
	_ = "STUB: not implemented"
	return *new(duckprovisionedserviceable.ProvisionedServicable), nil
}

func bindingUsesForbiddenSecret(forbiddenSecrets map[string]struct{}, secretRef *corev1.LocalObjectReference) bool {
	_ = "STUB: not implemented"
	return false
}

func (g *Generator) fetchServiceAccounts(ctx context.Context, build BuildPodable) ([]corev1.ServiceAccount, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *Generator) fetchBuildSecrets(ctx context.Context, build BuildPodable) ([]corev1.Secret, []corev1.LocalObjectReference, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (g *Generator) fetchBuilderConfig(ctx context.Context, build BuildPodable) (buildapi.BuildPodBuilderConfig, error) {
	_ = "STUB: not implemented"
	return *new(buildapi.BuildPodBuilderConfig), nil
}

func (g *Generator) resolveBuilderRef(ctx context.Context, ref *corev1.ObjectReference) (string, authn.Keychain, error) {
	_ = "STUB: not implemented"
	return "", *new(authn.Keychain), nil
}

func (g *Generator) resolveBuilderImage(ctx context.Context, build BuildPodable) (string, authn.Keychain, error) {
	_ = "STUB: not implemented"
	return "", *new(authn.Keychain), nil
}

func (g *Generator) resolveBuilder(ctx context.Context, build BuildPodable) (string, authn.Keychain, error) {
	_ = "STUB: not implemented"
	return "", *new(authn.Keychain), nil
}

func parseCNBID(image ggcrv1.Image, env string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
