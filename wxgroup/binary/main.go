package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// 定义一个 32 位整数
	value := uint32(0x12345678)

	// 将整数转换为大端字节序的字节数组
	var bigEndianBuffer bytes.Buffer
	binary.Write(&bigEndianBuffer, binary.BigEndian, value)

	// 将整数转换为小端字节序的字节数组
	var littleEndianBuffer bytes.Buffer
	binary.Write(&littleEndianBuffer, binary.LittleEndian, value)

	// 输出大端字节序的字节数组
	fmt.Println("Big Endian:", bigEndianBuffer.Bytes())

	// 输出小端字节序的字节数组
	fmt.Println("Little Endian:", littleEndianBuffer.Bytes())

	// 读取大端字节序的字节数组
	var readValueBigEndian uint32
	err := binary.Read(&bigEndianBuffer, binary.BigEndian, &readValueBigEndian)
	if err != nil {
		fmt.Println("Error reading big endian:", err)
		return
	}
	fmt.Println("Read Big Endian Value:", readValueBigEndian)

	// 读取小端字节序的字节数组
	var readValueLittleEndian uint32
	err = binary.Read(&littleEndianBuffer, binary.LittleEndian, &readValueLittleEndian)
	if err != nil {
		fmt.Println("Error reading little endian:", err)
		return
	}
	fmt.Println("Read Little Endian Value:", readValueLittleEndian)
}
