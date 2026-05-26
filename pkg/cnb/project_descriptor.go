package cnb

import (
	"log"
)

const defaultProjectDescriptorPath = "project.toml"

func ProcessProjectDescriptor(appDir, descriptorPath, platformDir string, logger *log.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

func parseProjectDescriptor(file string, logger *log.Logger) (descriptorV2, error) {
	_ = "STUB: not implemented"
	return *new(descriptorV2), nil
}

// v1 descriptor

func v1ToV2(v1 descriptorV1) descriptorV2 { _ = "STUB: not implemented"; return *new(descriptorV2) }

func processFiles(appDir string, d build) error { _ = "STUB: not implemented"; return nil }

// We only want to remove paths that don't match the patterns and are
// files otherwise we will end up removing too much.
// For eg if the include = ["*jar"]
// All the directories will not match the pattern and hence be removed.
// On the other hand if a directory is excluded/included,
// for eg include = "my-dir" files under "my-dir" will match the pattern and not be removed.

func getFileFilter(d build) (func(string) bool, error) { _ = "STUB: not implemented"; return nil, nil }

type descriptorV2 struct {
	Project project `toml:"_"`
	IO      ioTable `toml:"io"`
}

// Because CNB Project Descriptor v0.2 has two ways for defining environment variables.
// see https://github.com/buildpacks/spec/blob/main/extensions/project-descriptor.md#iobuildpacksbuildenv-optional
// This function calculates the final environment variables
func (d *descriptorV2) env() []envVariable { _ = "STUB: not implemented"; return nil }

type project struct {
	SchemaVersion string `toml:"schema-version"`
}

type ioTable struct {
	Buildpacks cnbTableV2 `toml:"buildpacks"`
}

type cnbTableV2 struct {
	build              `toml:",inline"`
	buildEnvVariableV2 `toml:",inline"`
	Group              []buildpack `toml:"group"`
}

type build struct {
	Include []string `toml:"include"`
	Exclude []string `toml:"exclude"`
	Builder string   `toml:"builder"`
}

type buildEnvVariableV2 struct {
	BuildEnv buildEnvVariable `toml:"build"`
	// Deprecated: use `[[io.buildpacks.build.env]]` instead. see https://github.com/buildpacks/pack/pull/1479
	EnvBuild envBuildVariable `toml:"env"`
}

type buildpack struct {
	Id      string `json:"id" toml:"id"`
	Version string `json:"version" toml:"version"`
	Uri     string `json:"uri" toml:"uri"`
}

type descriptorV1 struct {
	Build cnbTableV1 `toml:"build"`
}

type cnbTableV1 struct {
	build      `toml:",inline"`
	Buildpacks []buildpack   `toml:"buildpacks"`
	Env        []envVariable `toml:"env"`
}
