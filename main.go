package main

import (
	"aaagame/tests/course/binarytest"
	"aaagame/tests/course/buildtags"
)

func main() {
	buildtags.Print()
}

func callTcp() {
	go binarytest.TCPServer()
	binarytest.Client()
}
