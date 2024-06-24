package interfaceassert

import "testing"

func TestInterfaceExample(t *testing.T) {
	InterfaceExample()
}

func TestInterfaceType(t *testing.T) {
	InterfaceType()
}

func BenchmarkLoopInt8Slice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopInt8Slice()
	}
}

func BenchmarkLoopIntSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopInt64Slice()
	}
}
