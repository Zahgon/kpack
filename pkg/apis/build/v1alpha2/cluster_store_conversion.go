package v1alpha2

import (
	"context"

	"knative.dev/pkg/apis"

	"github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
)

const (
	clusterStoreServiceAccountRefAnnotation = "kpack.io/clusterStoreServiceAccountRef"
)

func (s *ClusterStore) ConvertTo(_ context.Context, to apis.Convertible) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs *ClusterStoreSpec) convertToAnnotations(toAnnotations map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ClusterStore) ConvertFrom(_ context.Context, from apis.Convertible) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs *ClusterStoreSpec) convertTo(to *v1alpha1.ClusterStoreSpec) {
	_ = "STUB: not implemented"
	return
}

func (cs *ClusterStoreSpec) convertFrom(from *v1alpha1.ClusterStoreSpec) {
	_ = "STUB: not implemented"
	return
}

func (ct *ClusterStoreStatus) convertTo(to *v1alpha1.ClusterStoreStatus) {
	_ = "STUB: not implemented"
	return
}

func (ct *ClusterStoreStatus) convertFrom(from *v1alpha1.ClusterStoreStatus) {
	_ = "STUB: not implemented"
	return
}

func (s *ClusterStore) convertFromAnnotations(fromAnnotations *map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}
