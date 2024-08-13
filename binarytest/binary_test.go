package binarytest

import "testing"

func TestTCPServer(t *testing.T) {
	//Binary()
	go TCPServer()
	Client()
}

func TestBinary(t *testing.T) {
	Binary()
}
