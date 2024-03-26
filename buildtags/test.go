//go:build test1
// +build test1

package buildtags

import "fmt"

func init() {
	strings = append(strings, "mysql test")
}

type Hello struct {
	ID      int64
	Message string
}

func (h *Hello) String() string {
	return h.Message
}

func Say() {
	hi := Hello{ID: 1, Message: "Hello, I am test1 tag"}
	fmt.Printf("Say:%s\n", hi.String())
}
