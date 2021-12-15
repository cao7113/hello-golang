//go:build ignore

package buildtag

// Note: changed from "// +build" to "//go:build" since go1.17
// https://stackoverflow.com/questions/68360688/whats-the-difference-between-gobuild-and-build-directives

const IgnoredWord = "i-am-ignored"

func IgnoreString() string {
	return IgnoredWord
}
