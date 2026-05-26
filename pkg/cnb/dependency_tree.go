package cnb

import (
	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
)

type Node struct {
	Buildpack *corev1alpha1.BuildpackStatus
	Children  []*Node
}

// NewTree generates a list of dependency trees for the given buildpacks. A
// buildpack's dependency tree is just the buildpack's group[].order[] but in
// proper tree form.
func NewTree(buildpacks []corev1alpha1.BuildpackStatus) []*Node {
	_ = "STUB: not implemented"
	return nil
}

// explictly create a new var here, since using pointers in for-loops gets nasty

func makeTree(lookup map[string]*corev1alpha1.BuildpackStatus, id string) *Node {
	_ = "STUB: not implemented"
	return nil
}
