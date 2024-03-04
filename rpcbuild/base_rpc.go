package rpcbuild

import (
	"fmt"
	"net"
	"net/rpc"
)

// 基础版

type HelloService struct{}

func (s *HelloService) Say(req string, resp *string) error {
	*resp = "Hello: " + req
	return nil
}

func BaseRun() {
	err := rpc.RegisterName("HelloService", &HelloService{}) // 注意：name 作为客户端请求的服务器名
	if err != nil {
		exit("rpc register", err)
	}
	listener, err := net.Listen("tcp", ListenAddr)
	if err != nil {
		exit("rpc listen", err)
	}
	defer listener.Close()
	conn, err := listener.Accept()
	if err != nil {
		exit("rpc accept", err)
	}
	rpc.ServeConn(conn)
}

func BaseDial() {
	client, err := rpc.Dial("tcp", ListenAddr)
	if err != nil {
		exit("rpc dial", err)
	}
	defer client.Close()
	var reply string
	err = client.Call("HelloService.Say", "Hello", &reply)
	if err != nil {
		exit("rpc call", err)
	}
	fmt.Println(reply)
}
