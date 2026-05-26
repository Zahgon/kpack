package dockercreds

import (
	"github.com/google/go-containerregistry/pkg/authn"
)

const (
	SecretFilePathEnv = "CREDENTIAL_PROVIDER_SECRET_PATH"
)

func NewVolumeSecretKeychain() (authn.Keychain, error) {
	_ = "STUB: not implemented"
	return *new(authn.Keychain), nil
}
