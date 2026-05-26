package cosign

import (
	"context"

	"github.com/buildpacks/lifecycle/platform/files"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	cosignoptions "github.com/sigstore/cosign/v2/cmd/cosign/cli/options"
	cosignremote "github.com/sigstore/cosign/v2/pkg/oci/remote"
	corev1 "k8s.io/api/core/v1"

	"github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
)

const (
	CosignRepositoryEnv       = "COSIGN_REPOSITORY"
	CosignDockerMediaTypesEnv = "COSIGN_DOCKER_MEDIA_TYPES"
)

type SignFunc func(*cosignoptions.RootOptions, cosignoptions.KeyOpts, cosignoptions.SignOptions, []string) error

type FetchSignatureFunc func(name.Reference, ...cosignremote.Option) (name.Tag, error)

type BuilderSigner interface {
	SignBuilder(context.Context, string, []*corev1.Secret, authn.Keychain) ([]v1alpha2.CosignSignature, error)
}

type ImageSigner struct {
	signFunc           SignFunc
	fetchSignatureFunc FetchSignatureFunc
}

func NewImageSigner(signFunc SignFunc, fetchSignatureFunc FetchSignatureFunc) *ImageSigner {
	_ = "STUB: not implemented"
	return nil
}

func (s *ImageSigner) Sign(ro *cosignoptions.RootOptions, report files.Report, secretLocation string, annotations, cosignRepositories, cosignDockerMediaTypes map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ImageSigner) sign(ro *cosignoptions.RootOptions, refImage, digest, secretLocation, cosignSecret string, annotations, cosignRepositories, cosignDockerMediaTypes map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// When password file is not available, default empty password is used

func (s *ImageSigner) SignBuilder(
	ctx context.Context,
	imageReference string,
	serviceAccountSecrets []*corev1.Secret,
	builderKeychain authn.Keychain,
) ([]v1alpha2.CosignSignature, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func findCosignSecrets(secretLocation string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
