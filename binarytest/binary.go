package binarytest

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

/*
大端和小端

什么是字节序？
字节（Byte）是计算机世界的计量单位，比如，一部电视剧是10G个字节（1GB）,一张图片是1K个字节（1KB）。这些数据量有多有少，大多数时候，一个字节肯定是装不下的，这个时候我们就要用到多字节。在使用多字节的存的时候，就会涉及到一个顺序问题。
在计算机中，字节序是指多字节数据在计算机内存中存储或者网络传输时各字节的存储顺序。有两种储存数据的方式：大端字节序（big endian）和小端字节序（little endian）。举例来说，数值0x2211使用两个字节储存：高位字节是0x22，低位字节是0x11。

什么是大端序和小端序？
大端字节序：高位字节在前，低位字节在后，这是人类读写数值的方法。简单来说，就是按照从低地址到高地址的顺序存放数据的高位字节到低位字节，就如同例子中的0X2211
小端字节序：低位字节在前，高位字节在后，就是按照从低地址到高地址的顺序存放据的低位字节到高位字节。即以0x1122形式储存。
*/

func Binary() {
	var i uint32 = 380
	iptr := (*[4]byte)(unsafe.Pointer(&i))
	for _, ptr := range iptr {
		fmt.Printf("%02X", ptr)
	}
	fmt.Println()

	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, i)
	for _, bin := range b {
		fmt.Printf("%02X", bin)
	}
	fmt.Println()

	binary.BigEndian.PutUint32(b, i)
	for _, bin := range b {
		fmt.Printf("%02X", bin)
	}
	fmt.Println()
}
