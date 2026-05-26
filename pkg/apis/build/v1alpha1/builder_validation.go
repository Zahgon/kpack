package v1alpha1

import (
	"context"

	v1 "k8s.io/api/core/v1"

	"knative.dev/pkg/apis"
)

func (cb *Builder) SetDefaults(context.Context) { _ = "STUB: not implemented"; return }

func (cb *Builder) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (s *BuilderSpec) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (s *NamespacedBuilderSpec) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func validateStack(stack v1.ObjectReference) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func validateStore(store v1.ObjectReference) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}
