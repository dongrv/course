package binarytest

import "testing"

func TestBinary(t *testing.T) {
	//Binary()
	go TCPServer()
	Client()
}
