package synctest

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 同步原语测试

// Mutex 互斥锁
func Mutex() {
	var (
		mu  sync.Mutex
		foo int
	)

	go func() {
		mu.Lock()
		fmt.Println("g1:获得锁")
		time.Sleep(3 * time.Second)
		foo = 1
		fmt.Println("g1:write foo", foo)
		fmt.Println("g1:解锁")
		mu.Unlock()
	}()

	go func() {
		time.Sleep(time.Second)
		mu.Lock()
		fmt.Println("g2:获得锁")
		fmt.Println("read foo", foo)
		fmt.Println("g2:解锁")
		mu.Unlock()
	}()

	time.Sleep(6 * time.Second)

}

// RWMutex 读写锁
func RWMutex() {
	var (
		mu  sync.RWMutex
		foo int
	)
	// 读+写
	go func() {
		mu.Lock()
		fmt.Println("g1:获得写锁")
		time.Sleep(3 * time.Second)
		foo = 1
		fmt.Println("g1:write foo", foo)
		fmt.Println("g1:解锁")
		mu.Unlock()
	}()

	go func() {
		time.Sleep(time.Second)
		mu.RLock()
		fmt.Println("g2:获得读锁")
		fmt.Println("read foo", foo)
		fmt.Println("g2:解锁")
		mu.RUnlock()
	}()

	fmt.Println()

	mu.Lock()
	foo = 2
	mu.Unlock()

	// 读+读
	go func() {
		mu.RLock()
		time.Sleep(6 * time.Second)
		fmt.Println("g3:获得读锁")
		fmt.Println("read foo", foo)
		fmt.Println("g3:解锁")
		mu.RUnlock()
	}()

	go func() {
		time.Sleep(5 * time.Second)
		mu.RLock()
		fmt.Println("g4:获得读锁")
		fmt.Println("read foo", foo)
		fmt.Println("g4:解锁")
		mu.RUnlock()
	}()

	time.Sleep(10 * time.Second)

}

// WaitGroup 同步多协程
func WaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer func() {
			wg.Done()
			fmt.Println("g1:done")
		}()
		time.Sleep(3 * time.Second)
	}()

	go func() {
		defer func() {
			wg.Done()
			fmt.Println("g2:done")
		}()
		time.Sleep(5 * time.Second)
	}()

	wg.Wait()
	fmt.Println("all goroutine exit")
}

var done bool

func read(key string, c *sync.Cond) {
	c.L.Lock()
	for !done {
		c.Wait()
	}
	fmt.Println(key, "read")
	c.L.Unlock()
}

func write(key string, c *sync.Cond) {
	fmt.Println(key, "writing")
	time.Sleep(time.Second)
	c.L.Lock()
	done = true
	c.L.Unlock()
	fmt.Println("wake all read goroutine")
	c.Broadcast()
}

// Cond 条件变量
func Cond() {
	cond := sync.NewCond(&sync.Mutex{})
	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)

	write("writer", cond)

	time.Sleep(5 * time.Second)
}

func Atomic() {
	var ui uint64
	if atomic.CompareAndSwapUint64(&ui, 0, 1) {
		fmt.Println("ui cas successfully")
	}
	atomic.AddUint64(&ui, 1)
	fmt.Println("ui add:", ui)
	atomic.SwapUint64(&ui, 9)
	fmt.Println("ui swap:", ui)

	go func() {
		var ui64 atomic.Uint64
		ui64.Store(100)
		fmt.Println("ui64:", ui64.Load())
		ui64.CompareAndSwap(100, 1000)
		fmt.Println("ui64:", ui64.Load())
		ui64.Add(1)
		fmt.Println("ui64:", ui64.Load())
		ui64.Swap(20000)
		fmt.Println("ui64:", ui64.Load())
	}()

}

var once sync.Once

func Once() {
	var o = 1000
	fn := func() {
		o *= 1000
	}
	once.Do(fn)
	fmt.Println(o)

	once.Do(fn)
	fmt.Println(o)

}

func Map() {
	var m sync.Map
	m.Store("key1", "value1")
	if v, ok := m.Load("key1"); ok {
		fmt.Println("key1", v)
	}
	if v, ok := m.Load("key2"); ok {
		fmt.Println("key2:", v)
	}
	if m.CompareAndSwap("key1", "value1", "value2") {
		v, _ := m.Load("key1")
		fmt.Println("key1 cas successfully", v)
	}

	v, ok := m.LoadOrStore("key1", "value3")
	fmt.Println("key1 LoadOrStore", v, ok)
	vv, okk := m.LoadOrStore("key2", "value4")
	fmt.Println("key1 LoadOrStore2", vv, okk)

	if v2, ok2 := m.LoadAndDelete("key1"); ok2 {
		fmt.Println("key1 loaded and deleted", v2)
	}
}

var ready bool

func SyncCond() {
	fn := func() bool {
		return ready
	}
	cond := sync.NewCond(&sync.Mutex{})

	go func() {
		time.Sleep(time.Second)
		cond.L.Lock()
		ready = true
		cond.L.Unlock()
		cond.Signal()
		fmt.Println("signal 1")
		time.Sleep(time.Second)
		cond.Signal()
		fmt.Println("signal 2")
	}()

	go func() {
		cond.L.Lock()
		if !fn() {
			cond.Wait()
		}
		fmt.Println("done1")
		cond.L.Unlock()
	}()

	go func() {
		cond.L.Lock()
		if !fn() {
			cond.Wait()
		}
		fmt.Println("done2")
		cond.L.Unlock()
	}()

	time.Sleep(10 * time.Second)

}
func CondBlockMutex() {
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)
	start := time.Now()
	// 模拟资源
	dataReady := false

	go func() {
		cond.L.Lock() // 优先拿到锁
		defer cond.L.Unlock()

		// 模拟一段时间后数据准备完成
		time.Sleep(10 * time.Second)
		dataReady = true

		// 数据准备完成后发出信号
		cond.Signal()
		fmt.Println("return", time.Since(start))
	}()

	time.Sleep(time.Second)
	// 主 Goroutine 等待数据准备好
	cond.L.Lock() // 堵塞
	fmt.Println("here", time.Since(start))
	for !dataReady { // 此时不会进入判断逻辑里，dataReady == true
		cond.Wait()
	}
	// 此处的数据已准备就绪，可以进行后续处理
	fmt.Println("Data is ready now.")
	cond.L.Unlock()
}

// 测试接口断言

type Interface interface {
	Call()
}

type Demo struct {
	Name string
}

func (d *Demo) Call() {}

func (d *Demo) GetTs() int64 { return 1 }

type Interface2 interface{ GetTs() int64 }

func Transport(i Interface) {
	if v, ok := i.(Interface2); ok {
		fmt.Println(v.GetTs())
		return
	}
	fmt.Println("failed", i)
}

func InterfaceExample() {
	var i Interface
	i = &Demo{Name: "example"}
	Transport(i)
}
