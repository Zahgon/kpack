package cnb

import (
	"github.com/Masterminds/semver/v3"
)

var anyStackMinimumVersion = semver.MustParse("0.5")

func (bl BuildpackLayerInfo) supports(buildpackApis []string, id string, mixins []string, relaxedMixinContract bool) error {
	_ = "STUB: not implemented"
	return nil
}

//ignore meta-buildpacks

// as of buildpack API 0.10+ stacks are optional (and deprecated)

func validateRequiredMixins(providedMixins, requiredMixins []string, relaxedMixinContract bool) error {
	_ = "STUB: not implemented"
	return nil
}

func present(haystack []string, needle string) bool { _ = "STUB: not implemented"; return false }

func mixinPresent(mixins []string, mixin string, relaxedMixinContract bool) bool {
	_ = "STUB: not implemented"
	return false
}

// A buildpack's mixin requirements must be satisfied by the stack in one of the following scenarios.
// 1) the stack provides the mixin `run:<mixin>` and the buildpack requires `run:<mixin>`
// 2) the stack provides the mixin `build:<mixin>` and the buildpack requires `build:<mixin>`
// 3) the stack provides the mixin `<mixin>` and the buildpack requires `<mixin>`
// 4) the stack provides the mixin `<mixin>` and the buildpack requires `build:<mixin>`
// 5) the stack provides the mixin `<mixin>` and the buildpack requires `run:<mixin>`
// 6) the stack provides the mixin `<mixin>` and the buildpack requires both `run:<mixin>` and `build:<mixin>`
// 7) the stack provides the mixins `build:<mixin>` and `run:<mixin>` the buildpack requires `<mixin>`

func stageRemoved(needle string) string { _ = "STUB: not implemented"; return "" }

func isAnystack(stackId string, buildpackVersion *semver.Version) bool {
	_ = "STUB: not implemented"
	return false
}
