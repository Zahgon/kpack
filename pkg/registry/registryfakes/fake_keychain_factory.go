package registryfakes

import (
	"context"
	"testing"

	"github.com/google/go-containerregistry/pkg/authn"

	"github.com/pivotal/kpack/pkg/registry"
)

type keychainContainer struct {
	SecretRef registry.SecretRef
	Keychain  authn.Keychain
}

type FakeKeychainFactory struct {
	keychains []keychainContainer
}

func (f *FakeKeychainFactory) KeychainForSecretRef(ctx context.Context, secretRef registry.SecretRef) (authn.Keychain, error) {
	_ = "STUB: not implemented"
	return *new(authn.Keychain), nil
}

func (f *FakeKeychainFactory) AddKeychainForSecretRef(t *testing.T, secretRef registry.SecretRef, keychain authn.Keychain) {
	_ = "STUB: not implemented"
	return
}

func (f *FakeKeychainFactory) getKeychainForSecretRef(secretRef registry.SecretRef) (authn.Keychain, bool) {
	_ = "STUB: not implemented"
	return *new(authn.Keychain), false
}

type FakeKeychain struct {
	Name string
}

func (f *FakeKeychain) Resolve(authn.Resource) (authn.Authenticator, error) {
	_ = "STUB: not implemented"
	return *new(authn.Authenticator), nil
}
