package secret

import (
	corev1 "k8s.io/api/core/v1"
)

type KeyType int

const (
	CosignKeyType KeyType = iota
	PKCS8KeyType
)

type SigningKey struct {
	SecretName string
	Key        []byte
	Password   []byte
	Type       KeyType
}

func FilterCosignSigningSecrets(secrets []*corev1.Secret) []*corev1.Secret {
	_ = "STUB: not implemented"
	return nil
}

func FilterAndExtractSLSASecrets(secrets []*corev1.Secret) ([]SigningKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func filterCosignSecrets(serviceAccountSecrets []*corev1.Secret, annotation string) []*corev1.Secret {
	_ = "STUB: not implemented"
	return nil
}

func filterPrivateKeySecrets(serviceAccountSecrets []*corev1.Secret, annotation string) []*corev1.Secret {
	_ = "STUB: not implemented"
	return nil
}

func extractAttestationKeyFromSecrets(cosignSecrets []*corev1.Secret, pkcs8Secrets []*corev1.Secret) ([]SigningKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getKey(secrets []*corev1.Secret, keyField, passField string, keyType KeyType) ([]SigningKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
