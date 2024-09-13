package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Arith int

func (t *Arith) Multiple(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

type Args struct {
	A, B int
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	log.Println("Starting RPC server on :1234...")
	log.Fatal(http.ListenAndServe(":1234", nil))
}
