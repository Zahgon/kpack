package registry

import (
	"context"

	"github.com/google/go-containerregistry/pkg/authn"
	v1 "k8s.io/api/core/v1"
)

type SecretRef struct {
	ServiceAccount   string
	Namespace        string
	ImagePullSecrets []v1.LocalObjectReference
}

func (s SecretRef) IsNamespaced() bool { _ = "STUB: not implemented"; return false }

func (s SecretRef) ServiceAccountOrDefault() string { _ = "STUB: not implemented"; return "" }

type KeychainFactory interface {
	KeychainForSecretRef(context.Context, SecretRef) (authn.Keychain, error)
}
