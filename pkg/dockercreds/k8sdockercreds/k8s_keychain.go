package k8sdockercreds

import (
	"context"

	"github.com/google/go-containerregistry/pkg/authn"
	corev1 "k8s.io/api/core/v1"
	k8sclient "k8s.io/client-go/kubernetes"

	"github.com/pivotal/kpack/pkg/dockercreds/k8sdockercreds/azurecredentialhelperfix"
	"github.com/pivotal/kpack/pkg/registry"
	"github.com/pivotal/kpack/pkg/secret"
)

var azureFileKeychain = azurecredentialhelperfix.AzureFileKeychain() // To support AZURE_CONTAINER_REGISTRY_CONFIG

type k8sSecretKeychainFactory struct {
	client         k8sclient.Interface
	volumeKeychain authn.Keychain
}

func NewSecretKeychainFactory(client k8sclient.Interface) (registry.KeychainFactory, error) {
	_ = "STUB: not implemented"
	return *new(registry.KeychainFactory), nil
}

func (f *k8sSecretKeychainFactory) KeychainForSecretRef(ctx context.Context, ref registry.SecretRef) (authn.Keychain, error) {
	_ = "STUB: not implemented"
	return *new(authn.Keychain), nil
}

// k8s keychain with no secrets

func toStringPullSecrets(secrets []corev1.LocalObjectReference) []string {
	_ = "STUB: not implemented"
	return nil
}

func keychainFromServiceAccount(ctx context.Context, secretRef registry.SecretRef, fetcher *secret.Fetcher) (authn.Keychain, error) {
	_ = "STUB: not implemented"
	return *new(authn.Keychain), nil
}
