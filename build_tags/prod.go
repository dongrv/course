//go:build prod
// +build prod

package build_tags

func init() {
	strings = append(strings, "mysql prod")
}
