package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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

	// w 写入 pipe
	go func() {
		defer w.Close()
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

}
