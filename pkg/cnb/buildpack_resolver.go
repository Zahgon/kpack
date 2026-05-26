package cnb

import (
	"github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	v1 "k8s.io/api/core/v1"
)

// BuildpackResolver will attempt to resolve a Buildpack reference to a
// Buildpack from either the ClusterStore, Buildpacks, or ClusterBuildpacks
type BuildpackResolver interface {
	resolve(ref v1alpha2.BuilderBuildpackRef) (K8sRemoteBuildpack, error)
	ClusterStoreObservedGeneration() int64
}

type buildpackResolver struct {
	clusterstore      *v1alpha2.ClusterStore
	buildpacks        []*v1alpha2.Buildpack
	clusterBuildpacks []*v1alpha2.ClusterBuildpack
}

func NewBuildpackResolver(clusterStore *v1alpha2.ClusterStore, buildpacks []*v1alpha2.Buildpack, clusterBuildpacks []*v1alpha2.ClusterBuildpack) BuildpackResolver {
	_ = "STUB: not implemented"
	return *new(BuildpackResolver)
}

func (r *buildpackResolver) ClusterStoreObservedGeneration() int64 {
	_ = "STUB: not implemented"
	return 0
}

func (r *buildpackResolver) resolve(ref v1alpha2.BuilderBuildpackRef) (K8sRemoteBuildpack, error) {
	_ = "STUB: not implemented"
	return *new(K8sRemoteBuildpack), nil
}

// TODO(chenbh):

func (r *buildpackResolver) resolveFromBuildpack(id string, buildpacks []*v1alpha2.Buildpack) ([]K8sRemoteBuildpack, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *buildpackResolver) resolveFromClusterBuildpack(id string, clusterBuildpacks []*v1alpha2.ClusterBuildpack) ([]K8sRemoteBuildpack, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *buildpackResolver) resolveFromClusterStore(id string, store *v1alpha2.ClusterStore) ([]K8sRemoteBuildpack, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// resolveFromObjectReference will get the object and figure out the root
// buildpack by converting it to a buildpack dependency tree
func (r *buildpackResolver) resolveFromObjectReference(ref v1.ObjectReference) (K8sRemoteBuildpack, error) {
	_ = "STUB: not implemented"
	return *new(K8sRemoteBuildpack), nil
}

// TODO: combine findBuildpack and findClusterBuildpack into a single func
// if/when golang generics has support for field values
func findBuildpack(ref v1.ObjectReference, buildpacks []*v1alpha2.Buildpack) *v1alpha2.Buildpack {
	_ = "STUB: not implemented"
	return nil
}

func findClusterBuildpack(ref v1.ObjectReference, clusterBuildpacks []*v1alpha2.ClusterBuildpack) *v1alpha2.ClusterBuildpack {
	_ = "STUB: not implemented"
	return nil
}

// TODO: error if the highest version has multiple diff ids
func highestVersion(matchingBuildpacks []K8sRemoteBuildpack) (K8sRemoteBuildpack, error) {
	_ = "STUB: not implemented"
	return *new(K8sRemoteBuildpack), nil
}
