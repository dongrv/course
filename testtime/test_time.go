package testtime

import (
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"
	"time"
)

func Time() {
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			fmt.Println("time over")
			return
		}
	}
}

// TimeAfter
// debug cpu profiler NewTimer 占比很大
// time.After 和 time.AfterFunc 均有定时器泄露风险，尤其在遍历和用户量大的的时候使用，会反复创建 timer 且不会释放，
// 表现为服务器卡顿，在一定周期内CPU占用渐进式上升
func TimeAfter() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/debug/pprof/pl", pprof.Profile)

	go func() {
		if err = http.Serve(ln, nil); err != nil {
			return
		}
	}()

	n := make(chan int, 10000)
	i := 0
	go func() {
		for {
			i++
			n <- i
		}
	}()
	for {
		select {
		case <-time.After(500 * time.Millisecond):
		case <-n:
			continue
		}
	}

}
