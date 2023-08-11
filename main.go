package main

import "aaagame/tests/course/binarytest"

func main() {
	go binarytest.TCPServer()
	binarytest.Client()
}
