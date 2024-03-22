package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type myReader struct{}

func (r *myReader) Read(p []byte) (int, error) {
	n := copy(p, "hello")
	return n, nil
}

type myWriter struct {
	buf    []byte
	offset int
}

func (w *myWriter) Write(p []byte) (int, error) {
	w.buf = append(w.buf, p...)
	w.offset = len(w.buf) - 1
	fmt.Println(string(w.buf))
	return w.offset, nil
}

func readFile() {
	f, err := os.OpenFile("./test.txt", os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer f.Close() // 是否资源
	buf := make([]byte, 16)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

func readBuffer() {
	strReader := strings.NewReader("Hello, world!")
	buf := bufio.NewReader(strReader)
	readbf := make([]byte, 16)
	for {
		n, err := buf.Read(readbf)
		if err == io.EOF {
			return
		}
		fmt.Println(string(readbf[:n]))
	}
}

func seekOffset() {
	f, err := os.Open("./test.txt")
	if err != nil {
		return
	}
	defer f.Close() // 释放资源
	_, err = f.Seek(3, io.SeekStart)
	if err != nil {
		return
	}
	buf := make([]byte, 16)
	f.Read(buf)
	fmt.Println("seek:", string(buf))

}

func ioPipe() {
	r, w := io.Pipe()

	defer func() {
		r.Close()
		w.Close()
	}()

	// w 写入 pipe
	go func() {
		w.Write([]byte("hello world"))
	}()

	//  r 读取pipe
	buf := make([]byte, 16)
	n, err := r.Read(buf)
	if err != nil {
		return
	}
	fmt.Println("read pipe:", string(buf[:n]))
}

func ioCopy() {
	src, err := os.Open("./test.txt")
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.Create("./dsttest.txt")
	if err != nil {
		return
	}
	defer dst.Close()
	io.Copy(dst, src) // 复制src内容到dst

}

func multiReader() {
	r1 := strings.NewReader("hello")
	r2 := strings.NewReader(" world")
	r := io.MultiReader(r1, r2)
	buf := make([]byte, 16)
	r.Read(buf)
	fmt.Println("multireader:", string(buf))
}

func multiWriter() {
	// 标准输出
	w := io.MultiWriter(os.Stdout, os.Stderr)
	r := strings.NewReader("hello writer\n")
	_, err := io.Copy(w, r)
	if err != nil {
		fmt.Println("multiwriter err:", err.Error())
	}
	time.Sleep(time.Second)

	// buffer
	var b1, b2 bytes.Buffer
	b := io.MultiWriter(&b1, &b2)
	b.Write([]byte("hello buffer")) // or io.Copy(b,strings.NewReader("hello buffer"))
	fmt.Println("b1:", b1.String(), "b2:", b2.String())
}

type RW interface {
	io.Reader
	io.Writer
}

func readerWriter() {
	var rw RW = os.Stdin
	buf := make([]byte, 16)
	rw.Read(buf)
	fmt.Println("reader:", string(buf))
	rw.Write(buf) // TODO 直接写入没起作用，不是标准输入的原因？
	fmt.Println("reader:2", string(buf))
}

func main() {
	// Test io.Reader and io.Writer
	r := myReader{}
	w := myWriter{}
	p := make([]byte, 16)
	r.Read(p)
	w.Write(p)

	// Test file read
	readFile()

	// Test buffer read
	readBuffer()

	// Test file seek
	seekOffset()

	// Test io.Pipe
	ioPipe()

	// Test io.Copy
	ioCopy()

	// Test io.MultiReader
	multiReader()

	// Test io.MultiWriter
	multiWriter()

	// Test reader writer
	readerWriter()

}

/*

io 操作注意事项

1. 检查操作返回值
2. 关闭对象释放资源
3. 避免读取/写入过大数据
4. 清理无用缓冲区
5. 设置超时与截止时间


io 包展示了一些好的接口设计思想:

1. 面向接口而非实现编程
3. 通过组合扩展接口
3. 最小公共方法集
4. 解耦相关功能

io 操作也可以通过各种优化提升性能:

1. 使用缓存器
2. 预分配大小
3. 多路复用 IO 操作
4. 并发读写
5. 复用 Buffer
6. 优化 Seek 操作

*/
