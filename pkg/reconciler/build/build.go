package build

import (
	"context"

	"github.com/google/go-containerregistry/pkg/authn"
	ggcrv1 "github.com/google/go-containerregistry/pkg/v1"
	intoto "github.com/in-toto/in-toto-golang/in_toto"
	corev1 "k8s.io/api/core/v1"
	corev1Informers "k8s.io/client-go/informers/core/v1"
	k8sclient "k8s.io/client-go/kubernetes"
	v1Listers "k8s.io/client-go/listers/core/v1"
	"knative.dev/pkg/controller"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
	"github.com/pivotal/kpack/pkg/buildchange"
	"github.com/pivotal/kpack/pkg/buildpod"
	"github.com/pivotal/kpack/pkg/client/clientset/versioned"
	buildinformers "github.com/pivotal/kpack/pkg/client/informers/externalversions/build/v1alpha2"
	buildlisters "github.com/pivotal/kpack/pkg/client/listers/build/v1alpha2"
	"github.com/pivotal/kpack/pkg/cnb"
	"github.com/pivotal/kpack/pkg/config"
	"github.com/pivotal/kpack/pkg/reconciler"
	"github.com/pivotal/kpack/pkg/registry"
	"github.com/pivotal/kpack/pkg/slsa"
)

const (
	ReconcilerName  = "Builds"
	Kind            = "Build"
	k8sOSLabel      = "kubernetes.io/os"
	ReasonCompleted = "Completed"
)

//go:generate counterfeiter . MetadataRetriever
type MetadataRetriever interface {
	GetBuildMetadata(string, string, authn.Keychain) (*cnb.BuildMetadata, error)
}

type PodGenerator interface {
	Generate(context.Context, buildpod.BuildPodable) (*corev1.Pod, error)
}

type PodProgressLogger interface {
	GetTerminationMessage(pod *corev1.Pod, s *corev1.ContainerStatus) (string, error)
}

//go:generate counterfeiter . SLSAAttester
type SLSAAttester interface {
	AttestBuild(build *buildapi.Build, buildMetadata *cnb.BuildMetadata, pod *corev1.Pod, builderAndAppKeychain authn.Keychain, builderID slsa.BuilderID, depFns ...slsa.BuilderDependencyFn) (intoto.Statement, error)
	Sign(ctx context.Context, stmt intoto.Statement, signers ...slsa.Signer) ([]byte, error)
	Write(ctx context.Context, digestStr string, payload []byte, keychain authn.Keychain) (ggcrv1.Image, string, error)
}

//go:generate counterfeiter . SecretFetcher
type SecretFetcher interface {
	SecretsForServiceAccount(ctx context.Context, serviceAccount, namespace string) ([]*corev1.Secret, error)
	SecretsForSystemServiceAccount(context.Context) ([]*corev1.Secret, error)
}

func NewController(
	ctx context.Context, opt reconciler.Options, k8sClient k8sclient.Interface,
	informer buildinformers.BuildInformer, podInformer corev1Informers.PodInformer,
	metadataRetriever MetadataRetriever,
	podGenerator PodGenerator, podProgressLogger *buildchange.ProgressLogger,
	keychainFactory registry.KeychainFactory,
	attester SLSAAttester,
	secretFetcher SecretFetcher,
	featureFlags config.FeatureFlags,
) *controller.Impl {
	_ = "STUB: not implemented"
	return nil
}

type Reconciler struct {
	Client            versioned.Interface
	KeychainFactory   registry.KeychainFactory
	Lister            buildlisters.BuildLister
	MetadataRetriever MetadataRetriever
	K8sClient         k8sclient.Interface
	PodLister         v1Listers.PodLister
	PodGenerator      PodGenerator
	PodProgressLogger PodProgressLogger
	Attester          SLSAAttester
	SecretFetcher     SecretFetcher
	FeatureFlags      config.FeatureFlags
}

func (c *Reconciler) Reconcile(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Reconciler) reconcile(ctx context.Context, build *buildapi.Build) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Reconciler) setBuildReady(ctx context.Context, pod *corev1.Pod) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Reconciler) cleanupIfNeeded(ctx context.Context, pod *corev1.Pod) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	// check if pod is running && completion is terminated
	return nil, nil
}

func (c *Reconciler) reconcileBuildPod(ctx context.Context, build *buildapi.Build) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Reconciler) conditionForPod(pod *corev1.Pod, stepsCompleted []string) corev1alpha1.Conditions {
	_ = "STUB: not implemented"
	return *new(corev1alpha1.Conditions)
}

func stepStates(pod *corev1.Pod) []corev1.ContainerState { _ = "STUB: not implemented"; return nil }

func stepsCompleted(pod *corev1.Pod) []string { _ = "STUB: not implemented"; return nil }

func buildStepCompleted(s *corev1.ContainerStatus) bool { _ = "STUB: not implemented"; return false }

func (c *Reconciler) updateStatus(ctx context.Context, desired *buildapi.Build) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Reconciler) buildMetadataFromBuildPod(pod *corev1.Pod) (*cnb.BuildMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Reconciler) attestBuild(ctx context.Context, build *buildapi.Build, buildMetadata *cnb.BuildMetadata, pod *corev1.Pod) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Reconciler) attestBuildDeps(ctx context.Context, build *buildapi.Build, pod *corev1.Pod, secrets []*corev1.Secret) ([]slsa.BuilderDependencyFn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func contains(arr []string, s string) bool { _ = "STUB: not implemented"; return false }

func completionContainerExited(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func allContainersReady(pod *corev1.Pod) bool { _ = "STUB: not implemented"; return false }

func createContainerStateForBuild(s *corev1.ContainerStatus) corev1.ContainerState {
	_ = "STUB: not implemented"
	return *new(corev1.ContainerState)
}
