package v1alpha2

import (
	"context"

	"knative.dev/pkg/apis"

	"github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
)

func (b *Builder) ConvertTo(_ context.Context, to apis.Convertible) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Builder) ConvertFrom(_ context.Context, from apis.Convertible) error {
	_ = "STUB: not implemented"
	return nil
}

func (bs *NamespacedBuilderSpec) convertTo(to *v1alpha1.NamespacedBuilderSpec) {
	_ = "STUB: not implemented"
	return
}

func (bs *NamespacedBuilderSpec) convertFrom(from *v1alpha1.NamespacedBuilderSpec) {
	_ = "STUB: not implemented"
	return
}

func (bst *BuilderStatus) convertFrom(from *v1alpha1.BuilderStatus) {
	_ = "STUB: not implemented"
	return
}

func (bst *BuilderStatus) convertTo(to *v1alpha1.BuilderStatus) { _ = "STUB: not implemented"; return }
