package v1alpha2

import (
	"context"

	"knative.dev/pkg/apis"

	"github.com/pivotal/kpack/pkg/apis/build/v1alpha1"
)

func (b *Build) ConvertTo(_ context.Context, to apis.Convertible) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Build) ConvertFrom(_ context.Context, from apis.Convertible) error {
	_ = "STUB: not implemented"
	return nil
}

func (bs *BuildSpec) convertTo(to *v1alpha1.BuildSpec) { _ = "STUB: not implemented"; return }

func (bs *BuildSpec) convertFrom(from *v1alpha1.BuildSpec) { _ = "STUB: not implemented"; return }

func (bs *BuildStatus) convertFrom(from *v1alpha1.BuildStatus) { _ = "STUB: not implemented"; return }

func (bs *BuildStatus) convertTo(to *v1alpha1.BuildStatus) { _ = "STUB: not implemented"; return }
