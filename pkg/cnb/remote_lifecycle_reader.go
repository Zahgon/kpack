package cnb

import (
	"github.com/google/go-containerregistry/pkg/authn"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
)

const lifecycleBuilderMetadataLabel = "io.buildpacks.builder.metadata"

type RemoteLifecycleReader struct {
	RegistryClient RegistryClient
}

func (r *RemoteLifecycleReader) Read(keychain authn.Keychain, clusterLifecycleSpec buildapi.ClusterLifecycleSpec) (buildapi.ResolvedClusterLifecycle, error) {
	_ = "STUB: not implemented"
	return *new(buildapi.ResolvedClusterLifecycle), nil
}

func toBuildAPISet(from APISet) buildapi.APISet {
	_ = "STUB: not implemented"
	return *new(buildapi.APISet)
}
