package reconciler

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

type Key struct {
	GroupKind      schema.GroupKind
	NamespacedName types.NamespacedName
}

func (k Key) String() string { _ = "STUB: not implemented"; return "" }

func (k Key) WithNamespace(namespace string) Key { _ = "STUB: not implemented"; return *new(Key) }

type Object interface {
	GetName() string
	GetNamespace() string
	GetObjectKind() schema.ObjectKind
}

func KeyForObject(obj Object) Key { _ = "STUB: not implemented"; return *new(Key) }

type Tracker interface {
	Track(ref Key, obj types.NamespacedName)
	TrackKind(kind schema.GroupKind, obj types.NamespacedName)
	OnChanged(obj interface{})
}
