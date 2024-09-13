package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dial error:", err)
	}

	// 这里，这里?
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	args := &Args{8, 9}
	var reply int
	err = client.Call("Arith.Multiple", args, &reply)
	if err != nil {
		log.Fatal("Multiple error:", err)
	}
	fmt.Printf("Multiple: %d*%d=%d\n", args.A, args.B, reply)
}
