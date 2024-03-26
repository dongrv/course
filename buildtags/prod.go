//go:build prod
// +build prod

package buildtags

func init() {
	strings = append(strings, "mysql prod")
}
