package cnb

import (
	"github.com/google/go-containerregistry/pkg/authn"

	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
)

type RemoteBuildpackReader struct {
	RegistryClient RegistryClient
}

func (r *RemoteBuildpackReader) Read(keychain authn.Keychain, storeImages []corev1alpha1.ImageSource) ([]corev1alpha1.BuildpackStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
