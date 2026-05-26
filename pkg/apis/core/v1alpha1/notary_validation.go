package v1alpha1

import (
	"context"

	"knative.dev/pkg/apis"
)

func (n *NotaryConfig) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (n *NotaryV1Config) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}
