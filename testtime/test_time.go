package testtime

import (
	"fmt"
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

func TimeAfter() {
	stop := make(chan struct{}, 1)
	go func() {
		time.Sleep(60 * time.Second)
		stop <- struct{}{}
	}()

	for {
		select {
		case <-time.After(time.Millisecond):

		case <-stop:
			fmt.Println("time over time after")
			return
		}
	}
	// debug cpu profiler NewTimer 占比很大
	// time.After 和 time.AfterFunc 均有定时器泄露风险，尤其在遍历和用户量大的的时候使用，会反复创建 timer 且不会释放，
	// 表现为服务器卡顿，在一定周期内CPU占用渐进式上升
}
