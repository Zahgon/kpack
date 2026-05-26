package v1alpha2

import (
	"context"

	"knative.dev/pkg/apis"

	"github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
)

func (i *SourceResolver) ConvertTo(_ context.Context, to apis.Convertible) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *SourceResolver) ConvertFrom(_ context.Context, from apis.Convertible) error {
	_ = "STUB: not implemented"
	return nil
}

func (is *SourceResolverSpec) convertTo(to *v1alpha1.SourceResolverSpec) {
	_ = "STUB: not implemented"
	return
}

func (is *SourceResolverSpec) convertFrom(from *v1alpha1.SourceResolverSpec) {
	_ = "STUB: not implemented"
	return
}

func (is *SourceResolverStatus) convertFrom(from *v1alpha1.SourceResolverStatus) {
	_ = "STUB: not implemented"
	return
}

func (is *SourceResolverStatus) convertTo(to *v1alpha1.SourceResolverStatus) {
	_ = "STUB: not implemented"
	return
}
