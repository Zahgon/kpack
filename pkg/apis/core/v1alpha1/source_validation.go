package v1alpha1

import (
	"context"

	"knative.dev/pkg/apis"
)

func (s *SourceConfig) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (g *Git) Validate(ctx context.Context) *apis.FieldError { _ = "STUB: not implemented"; return nil }

func (b *Blob) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}
