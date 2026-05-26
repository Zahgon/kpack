package cnb

import (
	"github.com/google/go-containerregistry/pkg/authn"
	ggcrv1 "github.com/google/go-containerregistry/pkg/v1"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
)

const (
	MixinsLabel = "io.buildpacks.stack.mixins"
	StackLabel  = "io.buildpacks.stack.id"

	cnbUserId  = "CNB_USER_ID"
	cnbGroupId = "CNB_GROUP_ID"
)

type RemoteStackReader struct {
	RegistryClient RegistryClient
}

func (r *RemoteStackReader) Read(keychain authn.Keychain, clusterStackSpec buildapi.ClusterStackSpec) (buildapi.ResolvedClusterStack, error) {
	_ = "STUB: not implemented"
	return *new(buildapi.ResolvedClusterStack), nil
}

func validateStackId(stackId string, buildImage ggcrv1.Image, runImage ggcrv1.Image) error {
	_ = "STUB: not implemented"
	return nil
}

func readMixins(image ggcrv1.Image) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func mixins(build, run []string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func classifyMixins(mixins []string, validPrefix, invalidPrefix string) (valid []string, invalid []string, common []string) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func missingCommonRunMixins(build []string, run []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func parseCNBID(image ggcrv1.Image, env string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
