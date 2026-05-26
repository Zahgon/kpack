package v1alpha2

import (
	"context"

	"knative.dev/pkg/apis"

	"github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
)

func (b *ClusterBuilder) ConvertTo(_ context.Context, to apis.Convertible) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *ClusterBuilder) ConvertFrom(_ context.Context, from apis.Convertible) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs *ClusterBuilderSpec) convertTo(to *v1alpha1.ClusterBuilderSpec) {
	_ = "STUB: not implemented"
	return
}

func (cs *ClusterBuilderSpec) convertFrom(from *v1alpha1.ClusterBuilderSpec) {
	_ = "STUB: not implemented"
	return
}
