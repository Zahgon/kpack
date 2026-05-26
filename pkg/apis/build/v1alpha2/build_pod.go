package v1alpha2

import (
	"path/filepath"
	"time"

	"github.com/Masterminds/semver/v3"
	corev1 "k8s.io/api/core/v1"
)

const (
	PrepareContainerName    = "prepare"
	AnalyzeContainerName    = "analyze"
	DetectContainerName     = "detect"
	RestoreContainerName    = "restore"
	BuildContainerName      = "build"
	ExportContainerName     = "export"
	RebaseContainerName     = "rebase"
	CompletionContainerName = "completion"

	secretVolumeNameTemplate     = "secret-volume-%v"
	pullSecretVolumeNameTemplate = "pull-secret-volume-%v"

	completionTerminationMessagePath = "/tmp/termination-log"
	cosignDefaultSecretPath          = "/var/build-secrets/cosign/%s"
	defaultSecretPath                = "/var/build-secrets/%s"
	ReportTOMLPath                   = "/var/report/report.toml"

	BuildLabel = "kpack.io/build"
	k8sOSLabel = "kubernetes.io/os"

	cosignDockerMediaTypesAnnotationPrefix = "kpack.io/cosign.docker-media-types"
	cosignRespositoryAnnotationPrefix      = "kpack.io/cosign.repository"
	DOCKERSecretAnnotationPrefix           = "kpack.io/docker"
	GITSecretAnnotationPrefix              = "kpack.io/git"
	BlobSecretAnnotationPrefix             = "kpack.io/blob"
	IstioInject                            = "sidecar.istio.io/inject"
	BuildReadyAnnotation                   = "build.kpack.io/ready"

	cosignSecretDataCosignKey = "cosign.key"

	cacheVolumeName                     = "cache-dir"
	homeVolumeName                      = "home-dir"
	layersVolumeName                    = "layers-dir"
	networkWaitLauncherVolumeName       = "network-wait-launcher-dir"
	buildWaitVolumeName                 = "build-wait-dir"
	downwardVolumeName                  = "downward-api-dir"
	notaryVolumeName                    = "notary-dir"
	platformVolumeName                  = "platform-dir"
	registrySourcePullSecretsVolumeName = "registry-source-pull-secrets-dir"
	reportVolumeName                    = "report-dir"
	workspaceVolumeName                 = "workspace-dir"

	buildChangesEnvVar           = "BUILD_CHANGES"
	CacheTagEnvVar               = "CACHE_TAG"
	platformApiVersionEnvVarName = "CNB_PLATFORM_API"
	serviceBindingRootEnvVar     = "SERVICE_BINDING_ROOT"
	TerminationMessagePathEnvVar = "TERMINATION_MESSAGE_PATH"

	PlatformEnvVarPrefix = "PLATFORM_ENV_"
)

var (
	PrepareCommand    = "/cnb/process/build-init"
	AnalyzeCommand    = "/cnb/lifecycle/analyzer"
	DetectCommand     = "/cnb/lifecycle/detector"
	RestoreCommand    = "/cnb/lifecycle/restorer"
	BuildCommand      = "/cnb/lifecycle/builder"
	ExportCommand     = "/cnb/lifecycle/exporter"
	CompletionCommand = "/cnb/process/completion"
	RebaseCommand     = "/cnb/process/rebase"
)

type ServiceBinding interface {
	ServiceName() string
}

type BuildPodImages struct {
	BuildInitImage   string
	BuildWaiterImage string
	CompletionImage  string
	RebaseImage      string
}

// +k8s:deepcopy-gen=false
type BuildContext struct {
	BuildPodBuilderConfig     BuildPodBuilderConfig
	Secrets                   []corev1.Secret
	Bindings                  []ServiceBinding
	ImagePullSecrets          []corev1.LocalObjectReference
	MaximumPlatformApiVersion *semver.Version
	InjectedSidecarSupport    bool
	SSHTrustUnknownHost       bool
}

type BuildPodBuilderConfig struct {
	StackID       string
	RunImage      string
	Uid           int64
	Gid           int64
	PlatformAPIs  []string
	ResolvedImage string
}

var (
	sourceMount = corev1.VolumeMount{
		Name:      workspaceVolumeName,
		MountPath: "/workspace",
	}
	homeMount = corev1.VolumeMount{
		Name:      homeVolumeName,
		MountPath: "/builder/home",
	}
	platformMount = corev1.VolumeMount{
		Name:      platformVolumeName,
		MountPath: "/platform",
	}
	cacheMount = corev1.VolumeMount{
		Name:      cacheVolumeName,
		MountPath: "/cache",
	}
	layersMount = corev1.VolumeMount{
		Name:      layersVolumeName,
		MountPath: "/layers",
	}
	projectMetadataMount = corev1.VolumeMount{
		Name:      layersVolumeName,
		MountPath: "/projectMetadata",
	}
	registrySourcePullSecretsMount = corev1.VolumeMount{
		Name:      registrySourcePullSecretsVolumeName,
		MountPath: "/registrySourcePullSecrets",
		ReadOnly:  true,
	}
	notaryV1Mount = corev1.VolumeMount{
		Name:      notaryVolumeName,
		MountPath: "/var/notary/v1",
		ReadOnly:  true,
	}
	reportMount = corev1.VolumeMount{
		Name:      reportVolumeName,
		MountPath: "/var/report",
		ReadOnly:  false,
	}
	homeEnv = corev1.EnvVar{
		Name:  "HOME",
		Value: "/builder/home",
	}
	serviceBindingRootEnv = corev1.EnvVar{
		Name:  serviceBindingRootEnvVar,
		Value: filepath.Join(platformMount.MountPath, "bindings"),
	}
	buildWaitMount = corev1.VolumeMount{
		Name:      buildWaitVolumeName,
		MountPath: "/buildWait",
		ReadOnly:  false,
	}
	downwardMount = corev1.VolumeMount{
		Name:      downwardVolumeName,
		MountPath: "/downward",
	}
)

type stepModifier func(corev1.Container) corev1.Container

func (b *Build) BuildPod(images BuildPodImages, buildContext BuildContext) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// empty string is a nop

// If the build fails, don't restart it.

func boolPointer(b bool) *bool { _ = "STUB: not implemented"; return nil }

func containerSecurityContext() *corev1.SecurityContext { _ = "STUB: not implemented"; return nil }

func podSecurityContext(config BuildPodBuilderConfig) *corev1.PodSecurityContext {
	_ = "STUB: not implemented"
	return nil
}

func setUpBuildWaiter(container corev1.Container, waitFile string) corev1.Container {
	_ = "STUB: not implemented"
	return *new(corev1.Container)
}

func (b *Build) useStandardContainers(buildWaiterImage string, pod *corev1.Pod) *corev1.Pod {
	_ = "STUB: not implemented"
	return nil
}

func (b *Build) notarySecretVolume() corev1.Volume {
	_ = "STUB: not implemented"
	return *new(corev1.Volume)
}

func (b *Build) notaryArgs() []string { _ = "STUB: not implemented"; return nil }

func (b *Build) cosignArgs() []string { _ = "STUB: not implemented"; return nil }

func (b *Build) rebasePod(buildContext BuildContext, images BuildPodImages) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Build) cacheVolume() []corev1.Volume { _ = "STUB: not implemented"; return nil }

func buildSecrets(includeBlobSecrets bool) func(corev1.Secret) bool {
	_ = "STUB: not implemented"
	return nil
}

func gitSecrets(secret corev1.Secret) bool { _ = "STUB: not implemented"; return false }

func blobSecrets(includeBlobSecret bool, secret corev1.Secret) bool {
	_ = "STUB: not implemented"
	return false
}

func dockerSecrets(secret corev1.Secret) bool { _ = "STUB: not implemented"; return false }

func (b *Build) setupSecretVolumesAndArgs(secrets []corev1.Secret, filter func(secret corev1.Secret) bool) ([]corev1.Volume, []corev1.VolumeMount, []string) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// ignoring secret

func (b *Build) setupImagePullVolumes(secrets []corev1.LocalObjectReference) ([]corev1.Volume, []corev1.VolumeMount, []string) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (b *Build) setupCosignVolumes(secrets []corev1.Secret) ([]corev1.Volume, []corev1.VolumeMount, []string) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

var (
	supportedPlatformAPIVersions = []*semver.Version{semver.MustParse("0.9"), semver.MustParse("0.8"), semver.MustParse("0.7")}
)

func (bc BuildContext) highestSupportedPlatformAPI(b *Build) (*semver.Version, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setupBindingVolumesAndMounts(bindings []ServiceBinding) ([]corev1.Volume, []corev1.VolumeMount, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func args(args ...[]string) []string { _ = "STUB: not implemented"; return nil }

func a(args ...string) []string { _ = "STUB: not implemented"; return nil }

func volumes(volumes ...[]corev1.Volume) []corev1.Volume { _ = "STUB: not implemented"; return nil }

func volumeMounts(volumes ...[]corev1.VolumeMount) []corev1.VolumeMount {
	_ = "STUB: not implemented"
	return nil
}

func deduplicate(lists ...[]corev1.LocalObjectReference) []corev1.LocalObjectReference {
	_ = "STUB: not implemented"
	return nil
}

func steps(f func(step func(corev1.Container, ...stepModifier))) []corev1.Container {
	_ = "STUB: not implemented"
	return nil
}

func cosignSecretArgs(secret corev1.Secret) []string { _ = "STUB: not implemented"; return nil }

func envs(envs []corev1.EnvVar, envVars ...corev1.EnvVar) []corev1.EnvVar {
	_ = "STUB: not implemented"
	return nil
}

func parseTime(providedTime string) (*time.Time, error) { _ = "STUB: not implemented"; return nil, nil }
