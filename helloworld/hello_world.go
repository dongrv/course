package helloworld

import (
	"fmt"
)

type inHeap struct {
	i int
}

func HelloWorld() {
	var a, b, c int         // 栈区
	var s []byte            // 栈区
	var str = "Hello World" // 常量区
	var p *inHeap           // 栈区
	p = new(inHeap)         // p本身在栈区，但新分配的内存在堆区
	fmt.Println(a, b, c, s, str, p)
}

func Equal1(id int32) bool {
	return id^1001 == 0 || id^1002 == 0 || id^1003 == 0 || id^1004 == 0 || id^1005 == 0 || id^1006 == 0 || id^1007 == 0 || id^1008 == 0
}

func Equal2(id int32) bool {
	return id == 1001 || id == 1002 || id == 1003 || id == 1004 || id == 1005 || id == 1006 || id == 1007 || id == 1008
}

func Slice() {
	a := []int{1}
	AppendSlice(a)
	fmt.Println(a)
}

func AppendSlice(s []int) {
	s[0] = 100
	s = append(s, 2, 3, 4, 5, 6, 7, 8)
}

func Bytes2Str(bytes []byte) (p string) {
	data := make([]byte, len(bytes))
	for i := 0; i < len(bytes); i++ {
		c := bytes[i]
		data[i] = c
	}
	// old
	//strHeader := (*reflect.StringHeader)(unsafe.Pointer(&p))
	//strHeader.Data = uintptr(unsafe.Pointer(&data[0]))
	//strHeader.Len = len(bytes)

	// new
	//p = unsafe.String(&data[0], len(bytes))

	return
}

type C struct {
	W *W
}

type W struct {
}

func (w *W) Write(p []byte) (n int, err error) {
	println("writing")
	return 0, nil
}

func TestDeliveryArgs(c *C) {
	println("w = nil")
	oc := c.W
	oc = nil
	_ = oc
}
func InterfaceDelivery() {
	w := &W{}
	c := &C{W: w}
	fmt.Printf("%v\n", c.W)
	TestDeliveryArgs(c)
	fmt.Printf("%v\n", c.W)
}
