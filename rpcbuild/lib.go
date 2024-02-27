package rpcbuild

import "os"

const ListenAddr = ":8080"

type Operator uint8

const (
	Plus     Operator = iota // 加
	Sub                      // 减
	Multiple                 // 乘
	Div                      // 除
)

func exit(from string, err error) {
	println(from, ": ", err.Error())
	os.Exit(-1)
}
