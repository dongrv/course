package rpcbuild

import (
	"testing"
	"time"
)

func TestCalculateRun(t *testing.T) {
	go CalculateRun()
	time.Sleep(100 * time.Millisecond)
	CalculateDial()
	select {}
}
