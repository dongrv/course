package testtcp

import (
	"fmt"
	"testing"
	"time"
)

func TestServe(t *testing.T) {
	println(fmt.Sprintf("测试模式:%d\n", Mode))
	go Serve()
	time.Sleep(time.Second)
	Dial()
}
