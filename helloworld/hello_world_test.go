package helloworld

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	HelloWorld()
}

func BenchmarkEqual1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Equal1(1009)
	}
	b.StopTimer()
}

func BenchmarkEqual2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Equal2(1009)
	}
	b.StopTimer()
}

func TestSlice(t *testing.T) {
	Slice()
}
