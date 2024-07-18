package readslice

import "testing"

func TestReadSlice(t *testing.T) {
	TestSlice()
}

func TestFields_Join(t *testing.T) {
	list := []string{"UID", "Name", "Age", "Address", "Country"}
	t.Log(Fields(list).Join())
}
