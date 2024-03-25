package interfaceassert

import (
	"fmt"
	"unsafe"
)

// 测试接口断言

type Interface interface {
	Call()
}

type Demo struct {
	Name string
}

func (d *Demo) Call() {}

func (d *Demo) GetTs() int64 { return 1 }

type Interface2 interface{ GetTs() int64 }

func Transport(i Interface) {
	if v, ok := i.(Interface2); ok {
		fmt.Println(v.GetTs())
		return
	}
	fmt.Println("failed", i)
}

func InterfaceExample() {
	var i Interface
	i = &Demo{Name: "example"}
	Transport(i)
}

func IsNil(i interface{}) bool {
	value := *(*uintptr)(unsafe.Pointer(&i))
	typ := *(*unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(&i)) + unsafe.Sizeof(value)))
	return typ == nil
}
