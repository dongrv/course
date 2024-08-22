package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello world!")
	}
}

func traceFunc() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	fmt.Println("Hello world!")
}

// go build -gcflags "-N"
// go tool objdump -S -s "main\.main" .\example.exe

// go tool trace .\trace.out
