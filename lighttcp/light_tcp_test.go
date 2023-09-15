package lighttcp

import "testing"

func TestGuard_Connect(t *testing.T) {
	g := &Guard{}
	g.Init()
	g.Connect()
}
