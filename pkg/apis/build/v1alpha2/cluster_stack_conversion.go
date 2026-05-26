package v1alpha2

import (
	"context"

	"knative.dev/pkg/apis"

	"github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
)

const (
	clusterStackServiceAccountRefAnnotation = "kpack.io/clusterStackServiceAccountRef"
)

func (s *ClusterStack) ConvertTo(_ context.Context, to apis.Convertible) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ClusterStack) ConvertFrom(_ context.Context, from apis.Convertible) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs *ClusterStackSpec) convertTo(to *v1alpha1.ClusterStackSpec) {
	_ = "STUB: not implemented"
	return
}

func (cs *ClusterStackSpec) convertToAnnotations(toAnnotations map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs *ClusterStackSpec) convertFrom(from *v1alpha1.ClusterStackSpec) {
	_ = "STUB: not implemented"
	return
}

func (ct *ClusterStackStatus) convertTo(to *v1alpha1.ClusterStackStatus) {
	_ = "STUB: not implemented"
	return
}

func (ct *ClusterStackStatus) convertFrom(from *v1alpha1.ClusterStackStatus) {
	_ = "STUB: not implemented"
	return
}

func (s *ClusterStack) convertFromAnnotations(fromAnnotations *map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}
