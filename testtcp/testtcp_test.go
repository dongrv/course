package testtcp

import (
	"testing"
	"time"
)

func TestServe(t *testing.T) {
	go Serve()
	time.Sleep(time.Second)
	Dial()
}
