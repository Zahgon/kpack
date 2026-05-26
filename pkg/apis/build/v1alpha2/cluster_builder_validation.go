package v1alpha2

import (
	"context"

	"knative.dev/pkg/apis"
)

func (ccb *ClusterBuilder) SetDefaults(context.Context) { _ = "STUB: not implemented"; return }

func (ccb *ClusterBuilder) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (ccbs *ClusterBuilderSpec) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}
