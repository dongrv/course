package interfaceassert

import "testing"

func TestInterfaceExample(t *testing.T) {
	InterfaceExample()
}

func TestInterfaceType(t *testing.T) {
	InterfaceType()
}

// go test -run TestMul/type -v

func TestMul(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		InterfaceExample()
	})
	t.Run("type", func(t *testing.T) {
		InterfaceType()
	})
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
