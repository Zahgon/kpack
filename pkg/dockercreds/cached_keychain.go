package dockercreds

import (
	"context"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/pivotal/kpack/pkg/registry"
)

type cacheKey string

type cachedKeychainFactory struct {
	keychainFactory registry.KeychainFactory
	cache           map[cacheKey]authn.Keychain
}

func NewCachedKeychainFactory(keychainFactory registry.KeychainFactory) registry.KeychainFactory {
	_ = "STUB: not implemented"
	return *new(registry.KeychainFactory)
}

func (f *cachedKeychainFactory) KeychainForSecretRef(ctx context.Context, secretRef registry.SecretRef) (authn.Keychain, error) {
	_ = "STUB: not implemented"
	return *new(authn.Keychain), nil
}

func makeKey(secretRef registry.SecretRef) cacheKey {
	_ = "STUB: not implemented"
	return *new(cacheKey)
}
