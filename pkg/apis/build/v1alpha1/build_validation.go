package v1alpha1

import (
	"context"

	"knative.dev/pkg/apis"
)

func (b *Build) SetDefaults(ctx context.Context) { _ = "STUB: not implemented"; return }

func (b *Build) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (bs *BuildSpec) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (bs *BuildSpec) validateImmutableFields(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (lb *LastBuild) Validate(context context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}
