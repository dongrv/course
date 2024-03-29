package readslice

import "fmt"

func TestSlice() {
	var msg = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	var data = []int{8, 9, 10}
	//fmt.Printf("%v\n", unsafe.Pointer(&msg))
	msg = msg[1:]
	fmt.Println(msg)
	//fmt.Printf("%v\n", unsafe.Pointer(&msg))
	msg = msg[:0]
	fmt.Println(msg, len(msg), cap(msg))
	//fmt.Printf("%v\n", unsafe.Pointer(&msg))
	msg = data[1:]
	fmt.Println(msg)
	//fmt.Printf("%v\n", unsafe.Pointer(&msg))

	// 扩容测试
	list := []int32{1, 2, 3}
	fmt.Printf("list-> len:%d cap:%d\n", len(list), cap(list))
	grow(list)
	fmt.Printf("grow list-> len:%d cap:%d\n", len(list), cap(list))
}

func grow(list []int32) {
	list = append(list, 1, 2, 3, 4, 5, 6) // TODO 触发扩容，创建了新引用类型，局部变量 list 指向 新引用地址，不再作用原引用
}
