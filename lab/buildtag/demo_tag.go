//go:build demo

package buildtag

// Note: changed from "// +build" to "//go:build" since go1.17
// https://stackoverflow.com/questions/68360688/whats-the-difference-between-gobuild-and-build-directives
// https://stackoverflow.com/questions/10646531/golang-conditional-compilation/67937234#67937234

func IsDemo() bool {
	return true
}
