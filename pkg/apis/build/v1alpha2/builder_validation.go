package v1alpha2

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

func validateOrder(order []BuilderOrderEntry) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func validateGroup(group BuilderOrderEntry) *apis.FieldError { _ = "STUB: not implemented"; return nil }

func validateBuildpackRef(ref BuilderBuildpackRef) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

// errs = errs.Also(validate.Image(ref.Image)).
// 	Also(apis.CheckDisallowedFields(ref.BuildpackInfo, v1alpha1.BuildpackInfo{})).
// 	Also(apis.CheckDisallowedFields(ref.ObjectReference, v1.ObjectReference{}))

func validateObjectRef(ref v1.ObjectReference, kinds []string) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}
