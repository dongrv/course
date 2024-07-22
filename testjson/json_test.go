package testjson

import "testing"

func TestUnmarshal(t *testing.T) {
	if err := Unmarshal(); err != nil {
		t.Error(err)
	}
}
