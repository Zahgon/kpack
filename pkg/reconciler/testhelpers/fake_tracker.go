package testhelpers

import (
	"github.com/pivotal/kpack/pkg/reconciler"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

type FakeTracker struct {
	objects map[string]map[types.NamespacedName]struct{}
	kinds   map[string]map[types.NamespacedName]struct{}
}

func (f *FakeTracker) Track(ref reconciler.Key, obj types.NamespacedName) {
	_ = "STUB: not implemented"
	return
}

func (f *FakeTracker) TrackKind(kind schema.GroupKind, obj types.NamespacedName) {
	_ = "STUB: not implemented"
	return
}

func (*FakeTracker) OnChanged(obj interface{}) { _ = "STUB: not implemented"; return }

func (f *FakeTracker) IsTracking(ref reconciler.Key, obj types.NamespacedName) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *FakeTracker) IsTrackingKind(kind schema.GroupKind, obj types.NamespacedName) bool {
	_ = "STUB: not implemented"
	return false
}

func (f FakeTracker) String() string { _ = "STUB: not implemented"; return "" }
