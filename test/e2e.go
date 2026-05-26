package test

import (
	"sync"
	"testing"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/pivotal/kpack/pkg/client/clientset/versioned"
)

var (
	setup         sync.Once
	client        *versioned.Clientset
	k8sClient     *kubernetes.Clientset
	dynamicClient dynamic.Interface
	clusterConfig *rest.Config
)

func newClients(t *testing.T) (*clients, error) { _ = "STUB: not implemented"; return nil, nil }

func getKubeConfig() string { _ = "STUB: not implemented"; return "" }

type clients struct {
	client        versioned.Interface
	k8sClient     kubernetes.Interface
	dynamicClient dynamic.Interface
}
