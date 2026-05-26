package v1alpha1

import (
	"context"
	"regexp"

	corev1 "k8s.io/api/core/v1"
	"knative.dev/pkg/apis"
)

func (bbs *BuildBuilderSpec) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func validateImage(value string) *apis.FieldError { _ = "STUB: not implemented"; return nil }

func validateBuilderRef(ref *corev1.ObjectReference) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func validateBuilderRefAndImageMutuallyExclusive(bbs *BuildBuilderSpec) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func validateBuilderOrImagePresent(bbs *BuildBuilderSpec) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (bs CNBBindings) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

// check name uniqueness

var bindingNameRE = regexp.MustCompile(`^[a-z0-9\-\.]{1,253}$`)

func (b *CNBBinding) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

// metadataRef is required

// secretRef is optional
