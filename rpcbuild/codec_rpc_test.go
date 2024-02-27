package rpcbuild

import (
	"testing"
	"time"
)

func TestCodecRun(t *testing.T) {
	go CodecRun()
	time.Sleep(100 * time.Millisecond)
	CodecDial()
	select {}
}
