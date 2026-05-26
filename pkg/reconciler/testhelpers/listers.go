package testhelpers

import (
	"k8s.io/apimachinery/pkg/runtime"
	fakekubeclientset "k8s.io/client-go/kubernetes/fake"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"knative.dev/pkg/reconciler/testing"

	"github.com/pivotal/kpack/pkg/client/clientset/versioned/fake"
	buildlisters "github.com/pivotal/kpack/pkg/client/listers/build/v1alpha2"
	"github.com/pivotal/kpack/pkg/duckbuilder"
)

var clientSetSchemes = []func(*runtime.Scheme) error{
	fake.AddToScheme,
	fakekubeclientset.AddToScheme,
}

type Listers struct {
	sorter testing.ObjectSorter
}

func NewListers(objs []runtime.Object) Listers { _ = "STUB: not implemented"; return *new(Listers) }

func (l *Listers) indexerFor(obj runtime.Object) cache.Indexer {
	_ = "STUB: not implemented"
	return *new(cache.Indexer)
}

func (l *Listers) BuildServiceObjects() []runtime.Object { _ = "STUB: not implemented"; return nil }

func (l *Listers) GetKubeObjects() []runtime.Object { _ = "STUB: not implemented"; return nil }

func (l *Listers) GetImageLister() buildlisters.ImageLister {
	_ = "STUB: not implemented"
	return *new(buildlisters.ImageLister)
}

func (l *Listers) GetBuildLister() buildlisters.BuildLister {
	_ = "STUB: not implemented"
	return *new(buildlisters.BuildLister)
}

func (l *Listers) GetBuilderLister() buildlisters.BuilderLister {
	_ = "STUB: not implemented"
	return *new(buildlisters.BuilderLister)
}

func (l *Listers) GetBuildpackLister() buildlisters.BuildpackLister {
	_ = "STUB: not implemented"
	return *new(buildlisters.BuildpackLister)
}

func (l *Listers) GetClusterBuilderLister() buildlisters.ClusterBuilderLister {
	_ = "STUB: not implemented"
	return *new(buildlisters.ClusterBuilderLister)
}

func (l *Listers) GetClusterBuildpackLister() buildlisters.ClusterBuildpackLister {
	_ = "STUB: not implemented"
	return *new(buildlisters.ClusterBuildpackLister)
}

func (l *Listers) GetClusterStoreLister() buildlisters.ClusterStoreLister {
	_ = "STUB: not implemented"
	return *new(buildlisters.ClusterStoreLister)
}

func (l *Listers) GetClusterStackLister() buildlisters.ClusterStackLister {
	_ = "STUB: not implemented"
	return *new(buildlisters.ClusterStackLister)
}

func (l *Listers) GetClusterLifecycleLister() buildlisters.ClusterLifecycleLister {
	_ = "STUB: not implemented"
	return *new(buildlisters.ClusterLifecycleLister)
}

func (l *Listers) GetSourceResolverLister() buildlisters.SourceResolverLister {
	_ = "STUB: not implemented"
	return *new(buildlisters.SourceResolverLister)
}

func (l *Listers) GetPersistentVolumeClaimLister() corev1listers.PersistentVolumeClaimLister {
	_ = "STUB: not implemented"
	return *new(corev1listers.PersistentVolumeClaimLister)
}

func (l *Listers) GetPodLister() corev1listers.PodLister {
	_ = "STUB: not implemented"
	return *new(corev1listers.PodLister)
}

func (l *Listers) GetConfigMapLister() corev1listers.ConfigMapLister {
	_ = "STUB: not implemented"
	return *new(corev1listers.ConfigMapLister)
}

func (l *Listers) GetDuckBuilderLister() *duckbuilder.DuckBuilderLister {
	_ = "STUB: not implemented"
	return nil
}
