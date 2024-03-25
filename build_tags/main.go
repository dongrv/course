package build_tags

import "fmt"

var strings []string

func Print() {
	for _, conf := range strings {
		fmt.Printf("conf:%s\n", conf)
	}
	Say()
}
