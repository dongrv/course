//go:build test1
// +build test1

package build_tags

func init() {
	strings = append(strings, "mysql test")
}
