package reconciler

import (
	"context"

	"knative.dev/pkg/controller"
)

type NetworkErrorReconciler struct {
	Reconciler controller.Reconciler
}

func (r *NetworkErrorReconciler) Reconcile(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// Re-queue the key if it's a network error.

type NetworkError struct {
	Err error
}

func (e *NetworkError) Error() string { _ = "STUB: not implemented"; return "" }
