package readinput

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// go doc fmt

func FmtScanln() {
	var (
		name string
		age  int
	)
	fmt.Println("Please input name and age, use space split:")
	_, _ = fmt.Scanln(&name, &age)
	fmt.Println("enter name:", name, " age:", age)
}

func ReadFmtInput() {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		// ctrl+d or ctrl+z trigger io.EOF error
		println(err.Error())
		return
	}
	fmt.Println("input:", input)
}

func ReadFmtInputs() {
	var a, b int
	_, err := fmt.Scan(&a, &b) // space spilt tow input variable
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println("a:", a, "b:", b)
}

func ReadBufioInput() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = scanner.Text()
		print("input:", input)
	}
	if err := scanner.Err(); err != nil {
		println(err.Error)
	}
}

func ReaderInput() {
	reader := strings.NewReader("Hello, world!")
	buf := make([]byte, 8)
	_, err := reader.Read(buf)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("reader:%v\n", buf)
}

// Sscan、Sscanf和Scanln
// 从字符串中扫描到数据到变量

// 使用 bufio 读取标准输入

func BufioReaderInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		println(err.Error())
		return
	}
	fmt.Printf("Your name:%s\n", input)
}
