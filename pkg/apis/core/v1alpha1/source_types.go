package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
)

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
type SourceConfig struct {
	Git      *Git      `json:"git,omitempty"`
	Blob     *Blob     `json:"blob,omitempty"`
	Registry *Registry `json:"registry,omitempty"`
	SubPath  string    `json:"subPath,omitempty"`
}

func (sc *SourceConfig) Source() Source { _ = "STUB: not implemented"; return *new(Source) }

type Source interface {
	BuildEnvVars() []corev1.EnvVar
	ImagePullSecretsVolume(name string) corev1.Volume
}

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
type Git struct {
	URL                  string `json:"url"`
	Revision             string `json:"revision"`
	InitializeSubmodules bool   `json:"initializeSubmodules,omitempty"`
}

func (g *Git) BuildEnvVars() []corev1.EnvVar { _ = "STUB: not implemented"; return nil }

func (in *Git) ImagePullSecretsVolume(name string) corev1.Volume {
	_ = "STUB: not implemented"
	return *new(corev1.Volume)
}

type BlobAuthKind string

const (
	BlobAuthNone   BlobAuthKind = ""
	BlobAuthHelper BlobAuthKind = "helper"
	BlobAuthSecret BlobAuthKind = "secret"
)

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
type Blob struct {
	URL             string `json:"url"`
	Auth            string `json:"auth,omitempty"`
	StripComponents int64  `json:"stripComponents,omitempty"`
}

func (b *Blob) ImagePullSecretsVolume(name string) corev1.Volume {
	_ = "STUB: not implemented"
	return *new(corev1.Volume)
}

func (b *Blob) BuildEnvVars() []corev1.EnvVar { _ = "STUB: not implemented"; return nil }

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
type Registry struct {
	Image string `json:"image"`
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +listType
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,15,rep,name=imagePullSecrets"`
}

func (r *Registry) ImagePullSecretsVolume(name string) corev1.Volume {
	_ = "STUB: not implemented"
	return *new(corev1.Volume)
}

func (r *Registry) BuildEnvVars() []corev1.EnvVar { _ = "STUB: not implemented"; return nil }

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
type ResolvedSourceConfig struct {
	Git      *ResolvedGitSource      `json:"git,omitempty"`
	Blob     *ResolvedBlobSource     `json:"blob,omitempty"`
	Registry *ResolvedRegistrySource `json:"registry,omitempty"`
}

func (sc ResolvedSourceConfig) ResolvedSource() ResolvedSource {
	_ = "STUB: not implemented"
	return *new(ResolvedSource)
}

type ResolvedSource interface {
	IsUnknown() bool
	IsPollable() bool
	SourceConfig() SourceConfig
}

type GitSourceKind string

const (
	Unknown GitSourceKind = "Unknown"
	Branch  GitSourceKind = "Branch"
	Tag     GitSourceKind = "Tag"
	Commit  GitSourceKind = "Commit"
)

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
type ResolvedGitSource struct {
	URL                  string        `json:"url"`
	Revision             string        `json:"revision"`
	SubPath              string        `json:"subPath,omitempty"`
	Tree                 string        `json:"tree,omitempty"`
	Type                 GitSourceKind `json:"type"`
	InitializeSubmodules bool          `json:"initializeSubmodules,omitempty"`
}

func (gs *ResolvedGitSource) SourceConfig() SourceConfig {
	_ = "STUB: not implemented"
	return *new(SourceConfig)
}

func (gs *ResolvedGitSource) IsUnknown() bool { _ = "STUB: not implemented"; return false }

func (gs *ResolvedGitSource) IsPollable() bool { _ = "STUB: not implemented"; return false }

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=true
type ResolvedBlobSource struct {
	URL             string `json:"url"`
	Auth            string `json:"auth,omitempty"`
	SubPath         string `json:"subPath,omitempty"`
	StripComponents int64  `json:"stripComponents,omitempty"`
}

func (bs *ResolvedBlobSource) SourceConfig() SourceConfig {
	_ = "STUB: not implemented"
	return *new(SourceConfig)
}

func (bs *ResolvedBlobSource) IsUnknown() bool { _ = "STUB: not implemented"; return false }

func (bs *ResolvedBlobSource) IsPollable() bool {
	_ = "STUB: not implemented"

	// +k8s:openapi-gen=true
	// +k8s:deepcopy-gen=true
	return false
}

type ResolvedRegistrySource struct {
	Image   string `json:"image"`
	SubPath string `json:"subPath,omitempty"`
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +listType
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,15,rep,name=imagePullSecrets"`
}

func (rs *ResolvedRegistrySource) SourceConfig() SourceConfig {
	_ = "STUB: not implemented"
	return *new(SourceConfig)
}

func (rs *ResolvedRegistrySource) IsUnknown() bool { _ = "STUB: not implemented"; return false }

func (rs *ResolvedRegistrySource) IsPollable() bool { _ = "STUB: not implemented"; return false }
