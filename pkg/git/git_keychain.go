package git

import (
	"net/url"

	"github.com/go-git/go-git/v5/plumbing/transport"

	"github.com/pivotal/kpack/pkg/secret"
)

type GitKeychain interface {
	Resolve(url string) (transport.AuthMethod, error)
}

type gitCredential interface {
	match(url *url.URL) bool
	auth() (transport.AuthMethod, error)
	name() string
}

type secretGitKeychain struct {
	creds []gitCredential
}

type gitSshAuthCred struct {
	fetchSecret          func() (secret.SSH, error)
	Domain               string
	SecretName           string
	sshTrustUnknownHosts bool
}

func (g gitSshAuthCred) auth() (transport.AuthMethod, error) {
	_ = "STUB: not implemented"
	return *new(transport.AuthMethod), nil
}

// the file is loaded when this callback is generated, so it's fine to remove the file right after

func (g gitSshAuthCred) match(url *url.URL) bool { _ = "STUB: not implemented"; return false }

func (g gitSshAuthCred) name() string { _ = "STUB: not implemented"; return "" }

type gitBasicAuthCred struct {
	fetchSecret func() (secret.BasicAuth, error)
	Domain      string
	SecretName  string
}

func (c gitBasicAuthCred) auth() (transport.AuthMethod, error) {
	_ = "STUB: not implemented"
	return *new(transport.AuthMethod), nil
}

func (c gitBasicAuthCred) match(url *url.URL) bool { _ = "STUB: not implemented"; return false }

func (c gitBasicAuthCred) name() string { _ = "STUB: not implemented"; return "" }

func NewMountedSecretGitKeychain(volumeName string, basicAuthSecrets, sshAuthSecrets []string, sshTrustUnknownHosts bool) (*secretGitKeychain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *secretGitKeychain) Resolve(rawUrl string) (transport.AuthMethod, error) {
	_ = "STUB: not implemented"
	return *new(transport.AuthMethod), nil
}
