package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type FakeProvisionedService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProvisionedServiceSpec   `json:"spec,omitempty"`
	Status ProvisionedServiceStatus `json:"status,omitempty"`
}

func (ps *FakeProvisionedService) DeepCopyObject() runtime.Object {
	_ = "STUB: not implemented"
	return *new(runtime.Object)
}

func (ps *FakeProvisionedService) GetGroupVersionKind() schema.GroupVersionKind {
	_ = "STUB: not implemented"
	return *new(schema.GroupVersionKind)
}

type ProvisionedServiceSpec struct {
	//random spec
	DatabaseSize string `json:"databaseSize,omitempty"`
}

type ProvisionedServiceStatus struct {
	Binding corev1.LocalObjectReference `json:"binding,omitempty"`
}
