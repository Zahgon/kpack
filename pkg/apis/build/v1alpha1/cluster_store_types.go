package v1alpha1

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"knative.dev/pkg/apis"

	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
)

const ClusterStoreKind = "ClusterStore"

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object,k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMetaAccessor

// +k8s:openapi-gen=true
type ClusterStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterStoreSpec   `json:"spec"`
	Status ClusterStoreStatus `json:"status"`
}

// +k8s:openapi-gen=true
type ClusterStoreSpec struct {
	// +listType
	Sources []corev1alpha1.ImageSource `json:"sources,omitempty"`
}

// +k8s:openapi-gen=true
type ClusterStoreStatus struct {
	corev1alpha1.Status `json:",inline"`

	// +listType
	Buildpacks []corev1alpha1.BuildpackStatus `json:"buildpacks,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +k8s:openapi-gen=true
type ClusterStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	// +k8s:listType=atomic
	Items []ClusterStore `json:"items"`
}

func (*ClusterStore) GetGroupVersionKind() schema.GroupVersionKind {
	_ = "STUB: not implemented"
	return *new(schema.GroupVersionKind)
}

func (s *ClusterStore) ConvertTo(_ context.Context, _ apis.Convertible) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ClusterStore) ConvertFrom(_ context.Context, _ apis.Convertible) error {
	_ = "STUB: not implemented"
	return nil
}
