//go:build dev
// +build dev

package build_tags

func init() {
	strings = append(strings, "mysql dev")
}
