package interfaceassert

import (
	"fmt"
	"runtime"
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

var store = map[interface{}]int{}

type Key struct {
}

func InterfaceType() {
	i := (*Key)(nil)
	store[i] = 1
	fmt.Printf("%d\n", len(store))
	delete(store, i)
	fmt.Printf("after %d\n", len(store))
}

func init() {
	runtime.GOMAXPROCS(1)
}

var (
	listInt  []int64
	listInt8 []int8
)

func init() {
	for i := 0; i < 1e3; i++ {
		listInt8 = append(listInt8, int8(i))
	}
	for i := 0; i < 1e3; i++ {
		listInt = append(listInt, int64(i))
	}
}

func LoopInt64Slice() {
	for _, i2 := range listInt {
		i2 = i2 * i2
	}
}

func LoopInt8Slice() {
	for _, i2 := range listInt8 {
		i2 = i2 * i2
	}
}
