package readslice

import "fmt"

func testSlice() {
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
}
