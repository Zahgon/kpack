package test

import (
	"context"
	"testing"
	"time"
)

func eventually(t *testing.T, fun func() bool, interval time.Duration, duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func printObject(t *testing.T, obj interface{}) { _ = "STUB: not implemented"; return }

func dumpK8s(t *testing.T, ctx context.Context, clients *clients, namespace string) {
	_ = "STUB: not implemented"
	return
}
