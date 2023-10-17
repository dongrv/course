package main

import (
	"aaagame/tests/course/binarytest"
	"unsafe"
)

func IsNil(i interface{}) bool {
	value := *(*uintptr)(unsafe.Pointer(&i))
	typ := *(*unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(&i)) + unsafe.Sizeof(value)))
	return typ == nil
}

func main() {
	go binarytest.TCPServer()
	binarytest.Client()
}
