package duckbuilder

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"

	buildapi "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	buildinformers "github.com/pivotal/kpack/pkg/client/informers/externalversions/build/v1alpha2"
	buildlisters "github.com/pivotal/kpack/pkg/client/listers/build/v1alpha2"
)

type DuckBuilderInformer struct {
	BuilderInformer        buildinformers.BuilderInformer
	ClusterBuilderInformer buildinformers.ClusterBuilderInformer
}

func (di *DuckBuilderInformer) AddBuilderEventHandler(handler cache.ResourceEventHandler) {
	_ = "STUB: not implemented"
	return
}

func (di *DuckBuilderInformer) AddClusterBuilderEventHandler(handler cache.ResourceEventHandler) {
	_ = "STUB: not implemented"
	return
}

func (di *DuckBuilderInformer) Lister() *DuckBuilderLister { _ = "STUB: not implemented"; return nil }

type DuckBuilderLister struct {
	BuilderLister        buildlisters.BuilderLister
	ClusterBuilderLister buildlisters.ClusterBuilderLister
}

func (bl *DuckBuilderLister) Namespace(namespace string) *DuckBuilderNamespaceLister {
	_ = "STUB: not implemented"
	return nil
}

type DuckBuilderNamespaceLister struct {
	DuckBuilderLister *DuckBuilderLister
	namespace         string
}

func (bl *DuckBuilderNamespaceLister) Get(reference corev1.ObjectReference) (*DuckBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertBuilder(builder *buildapi.Builder) *DuckBuilder { _ = "STUB: not implemented"; return nil }

func convertClusterBuilder(builder *buildapi.ClusterBuilder) *DuckBuilder {
	_ = "STUB: not implemented"
	return nil
}
