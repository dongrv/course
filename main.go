package main

import (
	"aaagame/tests/course/binarytest"
	"aaagame/tests/course/build_tags"
)

func main() {
	build_tags.Print()
}

func callTcp() {
	go binarytest.TCPServer()
	binarytest.Client()
}
