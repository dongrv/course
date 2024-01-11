package filter

import (
	"testing"
)

func TestFilter(t *testing.T) {
	f := &HashFilter{
		store:      make(map[unique]*metadata, 100),
		expiration: 10,
		frequency:  5,
		interval:   50,
	}
	d1 := f.Do("hello world %v test %d go %v build %+v")
	if d1.Push {
		t.Log(d1)
	}
	d2 := f.Do("hello world %f test %d go %v build %#v")
	if d2.Push {
		t.Log(d2)
	}

}
