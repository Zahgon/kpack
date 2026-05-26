package testing

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func GenerateFakeKeyPair(t *testing.T, secretName string, secretNamespace string, password string, annotations map[string]string) corev1.Secret {
	_ = "STUB: not implemented"
	return *new(corev1.Secret)
}

func Verify(t *testing.T, keyRef, imageRef string, annotations map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
