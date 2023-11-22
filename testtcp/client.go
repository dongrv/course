package testtcp

import (
	"encoding/json"
	"fmt"
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
		handleConn1(conn, "", "Client")
	case 2:
		go func() {
			time.Sleep(2 * time.Second)
			send := &HelloReq{Num: 1}
			msg, _ := json.Marshal(send)
			err := writeProcess(conn, string(msg), "Client")
			if err != nil {
				Exit(err)
			}
		}()
		handleConn2(conn, "", "Client")
	case 3:
		go readConn(conn, "", "Client")
		req := &HelloReq{Msg: "Hello"}
		msg, _ := json.Marshal(req)
		if err := writeMsgConn(conn, 1, msg, ""); err != nil {
			Exit(err)
		}
		fmt.Printf("Client send->%#v\n", req)
		time.Sleep(10 * time.Second)
		req2 := &HelloReq2{SumList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}
		msg, _ = json.Marshal(req2)
		if err := writeMsgConn(conn, 3, msg, ""); err != nil {
			Exit(err)
		}
		fmt.Printf("Client send2->%#v\n", req)
		time.Sleep(10 * time.Second)
	}
}
