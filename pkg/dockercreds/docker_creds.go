package dockercreds

import (
	"github.com/google/go-containerregistry/pkg/authn"
)

type DockerCreds map[string]authn.AuthConfig

func (c DockerCreds) Resolve(reg authn.Resource) (authn.Authenticator, error) {
	_ = "STUB: not implemented"
	return *new(authn.Authenticator), nil
}

// Fallback on anonymous.

func (c DockerCreds) Save(path string) error { _ = "STUB: not implemented"; return nil }

func (c DockerCreds) Append(a DockerCreds) (DockerCreds, error) {
	_ = "STUB: not implemented"
	return *new(DockerCreds), nil
}

func (c DockerCreds) contains(reg string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type dockerConfigJson struct {
	Auths DockerCreds `json:"auths"`
}
