//go:build dev
// +build dev

package buildtags

import "fmt"

func init() {
	strings = append(strings, "mysql dev")
}

type Hello struct {
	ID      int64
	Message string
}

func (h *Hello) String() string {
	return h.Message
}

func Say() {
	hi := Hello{ID: 1, Message: "Hello, I am dev tag"}
	fmt.Printf("Say:%s\n", hi.String())
}
