package generictype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwap(t *testing.T) {
	a, b := 1, 2
	c, d := Swap(a, b)
	if !assert.Equal(t, []interface{}{2, 1}, []interface{}{c, d}) {
		t.Failed()
	}

	// 实例化
	int1, int2 := Swap[int](10, 20)
	pair := NewPair[string]("Hello", "World")
	t.Log(int1, int2, pair)

}

func TestNewPair(t *testing.T) {
	t.Log(NewPair(1, 2))
}

func TestNewStack(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	t.Log(stack.Peek())

	for {
		meta, ok := stack.Pop()
		if !ok {
			break
		}
		t.Log(meta)
	}
}
