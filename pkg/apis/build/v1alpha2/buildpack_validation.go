package v1alpha2

import (
	"context"

	"knative.dev/pkg/apis"
)

func (cb *Buildpack) SetDefaults(context.Context) { _ = "STUB: not implemented"; return }

func (cb *Buildpack) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (s *BuildpackSpec) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}
