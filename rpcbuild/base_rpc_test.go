package rpcbuild

import (
	"testing"
	"time"
)

func TestRpcServer(t *testing.T) {
	go BaseRun()

	time.Sleep(10 * time.Millisecond)

	BaseDial()

	select {}
}
