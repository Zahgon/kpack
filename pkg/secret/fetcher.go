package secret

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	k8sclient "k8s.io/client-go/kubernetes"
)

type Fetcher struct {
	Client k8sclient.Interface

	SystemNamespace          string
	SystemServiceAccountName string
}

func (f *Fetcher) SecretsForServiceAccount(ctx context.Context, serviceAccount, namespace string) ([]*corev1.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *Fetcher) secretsFromServiceAccount(ctx context.Context, account *corev1.ServiceAccount, namespace string) ([]*corev1.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *Fetcher) SecretsForSystemServiceAccount(ctx context.Context) ([]*corev1.Secret, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
