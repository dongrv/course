package testtcp

import (
	"net"
	"os"
	"os/signal"
)

func Dial() {
	ln, err := net.Dial("tcp", addr)
	if err != nil {
		Exit(err)
	}
	println("The client is up")
	go handleRawConn(ln, "Client")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	select {
	case <-sig:
		break
	}
	_ = ln.Close()
	println("The client has exited")
}
