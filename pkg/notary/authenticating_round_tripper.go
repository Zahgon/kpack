package notary

import (
	"net/http"

	"github.com/google/go-containerregistry/pkg/authn"
)

type AuthenticatingRoundTripper struct {
	Token               string
	Keychain            authn.Keychain
	WrappedRoundTripper http.RoundTripper
}

func (a *AuthenticatingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractBearerOption(kind string, from string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type registryAuthResource struct {
	URL string
}

func parseRegistryAuthResource(realm string) (registryAuthResource, error) {
	_ = "STUB: not implemented"
	return *new(registryAuthResource), nil
}

func (r registryAuthResource) String() string { _ = "STUB: not implemented"; return "" }

func (r registryAuthResource) RegistryStr() string { _ = "STUB: not implemented"; return "" }
