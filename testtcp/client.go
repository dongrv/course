package testtcp

import (
	"encoding/json"
	"net"
	"os"
	"os/signal"
	"time"
)

func Dial() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		Exit(err)
	}
	println("The client is up")
	go sendMsg(conn)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	_ = conn.Close()
	println("The client has exited")
}

func sendMsg(conn net.Conn) {
	switch Mode {
	case 1:
		handleRawConn(conn, "", "Client")
	case 2:
		go func() {
			time.Sleep(2 * time.Second)
			send := &Send{Num: 1}
			msg, _ := json.Marshal(send)
			err := writeProcess(conn, string(msg), "Client")
			if err != nil {
				Exit(err)
			}
		}()
		handleProcessConn(conn, "", "Client")
	}
}
