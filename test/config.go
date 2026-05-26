package test

import (
	"testing"

	"github.com/google/go-containerregistry/pkg/authn"
	corev1 "k8s.io/api/core/v1"
)

type config struct {
	builder              string
	testRegistry         string
	testRegistryUsername string
	testRegistryPassword string
	gitSourcePrivateRepo string
	gitSourceUsername    string
	gitSourcePassword    string
	gitSourcePrivateKey  string
	imageTag             string
}

type dockerCredentials map[string]authn.AuthConfig

type dockerConfigJson struct {
	Auths dockerCredentials `json:"auths"`
}

const (
	lifecycleImage = "mirror.gcr.io/buildpacksio/lifecycle"
)

func loadConfig(t *testing.T) config { _ = "STUB: not implemented"; return *new(config) }

func (c *config) newImageTag() string { _ = "STUB: not implemented"; return "" }

func (c *config) makeRegistrySecret(secretName string, namespace string) (*corev1.Secret, error) {
	_ = "STUB: not implemented"
	return nil,

		// Handle path in registry
		nil
}

func (c *config) makeGitBasicAuthSecret(secretName, namespace string) (*corev1.Secret, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

// convert `github.com/org/repo` -> `https://github.com/org/repo.git`

func (c *config) makeGitSSHAuthSecret(secretName, namespace string) (*corev1.Secret, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

// convert `github.com/org/repo` -> `git@github.com:org/repo.git`
