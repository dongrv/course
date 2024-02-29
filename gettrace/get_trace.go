package gettrace

import (
	"os"
	"runtime"
	"runtime/trace"
)

func Trace() {

	runtime.GOMAXPROCS(1)

	file, err := os.OpenFile("./trace.out", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	trace.Start(file)
	defer trace.Stop()

	ch := make(chan string)
	go func() {
		ch <- "trace"
	}()
	<-ch

}
