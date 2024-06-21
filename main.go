package main

import (
	"aaagame/tests/course/binarytest"
	"aaagame/tests/course/ginx"
)

func main() {
	//buildtags.Print()
	ginx.Run()
}

func callTcp() {
	go binarytest.TCPServer()
	binarytest.Client()
}
