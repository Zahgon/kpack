package azurecredentialhelperfix

import (
	"sync"

	"github.com/google/go-containerregistry/pkg/authn"
	credentialprovider "github.com/vdemeester/k8s-pkg-credentialprovider"
)

var (
	once    sync.Once
	keyring credentialprovider.DockerKeyring
)

func AzureFileKeychain() authn.Keychain { _ = "STUB: not implemented"; return *new(authn.Keychain) }

// Copied from https://github.com/google/go-containerregistry/blob/d9bfbcb99e526b2a9417160e209b816e1b1fb6bd/pkg/authn/k8schain/k8schain.go#L141
type lazyProvider struct {
	kc    *keychain
	image string
}

// Authorization implements Authenticator.
func (lp lazyProvider) Authorization() (*authn.AuthConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type keychain struct {
	keyring credentialprovider.DockerKeyring
}

// Resolve implements authn.Keychain
func (kc *keychain) Resolve(target authn.Resource) (authn.Authenticator, error) {
	_ = "STUB: not implemented"
	return *new(authn.Authenticator), nil
}

// Lookup expects an image reference and we only have a registry.
