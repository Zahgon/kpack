package slsa

import (
	"time"

	"github.com/google/go-containerregistry/pkg/authn"
	intoto "github.com/in-toto/in-toto-golang/in_toto"
	slsav1 "github.com/in-toto/in-toto-golang/in_toto/slsa_provenance/v1"
	corev1 "k8s.io/api/core/v1"

	buildv1alpha2 "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	"github.com/pivotal/kpack/pkg/cnb"
	"github.com/pivotal/kpack/pkg/config"
)

type BuilderID string

const (
	SignedBuildID   BuilderID = "https://kpack.io/slsa/signed-build"
	UnsignedBuildID BuilderID = "https://kpack.io/slsa/unsigned-build"
	MediaTypeJSON             = "application/json"
)

type ImageReader interface {
	Read(keychain authn.Keychain, repoName string) (string, string, map[string]string, error)
}

type Attester struct {
	Version string

	ImageReader ImageReader

	Images   config.Images
	Features config.FeatureFlags
	Config   config.Config
}

func (a *Attester) AttestBuild(build *buildv1alpha2.Build, buildMetadata *cnb.BuildMetadata, pod *corev1.Pod, builderAndAppKeychain authn.Keychain, builderId BuilderID, depFns ...BuilderDependencyFn) (intoto.Statement, error) {
	_ = "STUB: not implemented"
	return *new(intoto.Statement), nil
}

type internalParams struct {
	BuilderImage string `json:"builderImage"`

	config.Config
	config.Images
	config.FeatureFlags
}

func (a *Attester) internalParamsFor(build *buildv1alpha2.Build) internalParams {
	_ = "STUB: not implemented"
	return *new(internalParams)
}

func getInvocationId(build *buildv1alpha2.Build, pod *corev1.Pod) string {
	_ = "STUB: not implemented"
	return ""
}

func getBuildType(version string) string { _ = "STUB: not implemented"; return "" }

func getStartStopTime(pod *corev1.Pod) (*time.Time, *time.Time) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertMap(orig map[string]string) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

type BuilderDependencyFn func() (slsav1.ResourceDescriptor, error)

type versionedObject struct {
	Name            string `json:"name"`
	ResourceVersion string `json:"resourceVersion"`
}

type K8sObject interface {
	GetName() string
	GetResourceVersion() string
}

// WithVersionedObject converts a kubernetes object to a SLSA ResourceDescriptor, where the name is
// the Kind, and the content is the json serialzed Name and ResourceVersion of the object.
func WithVersionedObject(kind string, obj K8sObject) BuilderDependencyFn {
	_ = "STUB: not implemented"
	return *new(BuilderDependencyFn)
}

// WithVersionedObjects is the same as WithVersionedObject but handles a slice of objects. These
// objects must have the same GVK
func WithVersionedObjects(kind string, objs []K8sObject) BuilderDependencyFn {
	_ = "STUB: not implemented"
	return *new(BuilderDependencyFn)
}
