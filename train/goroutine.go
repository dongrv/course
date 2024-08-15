package train

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func RunGoroutine() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			println("goroutine id is ", i)
		}(i)
	}
	wg.Wait()
}

// 自旋锁实现

type SpinLock struct {
	lock uint32
}

func NewSpinLock() *SpinLock {
	return &SpinLock{0}
}

func (sl *SpinLock) Lock() {
	for atomic.CompareAndSwapUint32(&sl.lock, 0, 1) == false {
		// 忙等待
	}
}

func (sl *SpinLock) Unlock() {
	atomic.StoreUint32(&sl.lock, 0)
}

func SpinRun() {
	spinLock := NewSpinLock()

	go func() {
		spinLock.Lock()
		time.Sleep(10 * time.Second) // 模拟长时间持有锁
		spinLock.Unlock()
	}()

	time.Sleep(time.Millisecond) // 给第一个协程一点时间先获取锁

	spinLock.Lock()
	defer spinLock.Unlock()
	fmt.Println("Critical section entered.")
}
