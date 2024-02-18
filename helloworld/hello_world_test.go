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

func TestBytes2Str(t *testing.T) {
	t.Log(Bytes2Str([]byte("hello world")))
}

func TestInterfaceDelivery(t *testing.T) {
	InterfaceDelivery()
}
