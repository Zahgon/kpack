package v1alpha2

import (
	"context"

	"knative.dev/pkg/apis"
)

func (cl *ClusterLifecycle) SetDefaults(context.Context) { _ = "STUB: not implemented"; return }

func (cl *ClusterLifecycle) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}

func (cls *ClusterLifecycleSpec) Validate(ctx context.Context) *apis.FieldError {
	_ = "STUB: not implemented"
	return nil
}
