package rpcbuild

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 增加JSON解码器

func CodecRun() {
	_ = RegisterService(&CalculateService{})
	listener, err := net.Listen("tcp", ListenAddr)
	if err != nil {
		exit("codec listen", err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			if ne, ok := err.(*net.OpError); ok && (ne.Temporary() && ne.Timeout()) {
				continue
			}
			exit("codec listen accept", err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) // 基于JSON编码的RPC服务
	}
}

func CodecDial() {
	ln, err := net.Dial("tcp", ListenAddr)
	if err != nil {
		exit("codec dial", err)
	}
	defer ln.Close()
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(ln))
	input := Input{Operator: Plus, Number1: 100, Number2: 100}
	var output float64
	<-client.Go(CalculateServiceName+".Calculate", input, &output, make(chan *rpc.Call, 1)).Done
	fmt.Printf("%s = %f\n", input.Expression(), output)
}
