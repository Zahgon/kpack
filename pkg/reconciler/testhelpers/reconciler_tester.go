package testhelpers

import (
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"
	"k8s.io/apimachinery/pkg/api/resource"
	rtesting "knative.dev/pkg/reconciler/testing"
)

func ReconcilerTester(t *testing.T, factory rtesting.Factory) SpecReconcilerTester {
	_ = "STUB: not implemented"
	return *new(SpecReconcilerTester)
}

type SpecReconcilerTester struct {
	t       *testing.T
	factory rtesting.Factory
}

func (rt SpecReconcilerTester) Test(test rtesting.TableRow) { _ = "STUB: not implemented"; return }

// Validate cached objects do not get soiled after controller loops

var (
	safeDeployDiff = cmpopts.IgnoreUnexported(resource.Quantity{})
)
