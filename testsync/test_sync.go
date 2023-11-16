package testsync

import (
	"fmt"
	"sync"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.Cond{L: &locker}

func Cond() {
	started := false
	for i := 0; i < 10; i++ {
		go func(i int) {
			cond.L.Lock()
			for !started {
				cond.Wait()
				fmt.Println(i)
			}
			cond.L.Unlock()
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Signal 1")
	locker.Lock()
	started = true
	locker.Unlock()
	cond.Signal()

	time.Sleep(time.Second)
	fmt.Println("Signal 2")
	cond.Broadcast()
	time.Sleep(time.Second)
}
