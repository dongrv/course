package filter

import (
	"testing"
)

func TestFilter(t *testing.T) {
	f := NewHashFilter(100, 50, 20)
	d1 := f.Do("hello world %v test %d go %v build %+v")
	if d1.Push {
		t.Log(d1)
	}
	d2 := f.Do("hello world %f test %d go %v build %#v")
	if d2.Push {
		t.Log(d2)
	}

}
