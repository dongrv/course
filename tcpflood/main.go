package main

import (
	"net"
	"sync"
	"sync/atomic"
	"time"
)

// const addr = "127.0.0.1:2001"
const addr = "192.168.1.6:2001"

var success, fail int64

func NewConn() (net.Conn, error) {
	return net.Dial("tcp", addr)
}

var emit bool

var conns []net.Conn

func Watch(wg *sync.WaitGroup, cond *sync.Cond) error {
	defer wg.Done()
	cond.L.Lock()
	for !emit {
		cond.Wait()
	}
	cond.L.Unlock()
	conn, err := NewConn()
	conns = append(conns, conn)
	return err
}

func Emit(cond *sync.Cond) {
	cond.L.Lock()
	emit = true
	cond.L.Unlock()
	cond.Broadcast()
	println("广播完成")
}

func main() {
	for i := 0; i < 10; i++ {
		{
			{
				var wg = &sync.WaitGroup{}

				cond := sync.NewCond(&sync.Mutex{})
				for i := 0; i < 512; i++ {
					wg.Add(1)
					go func() {
						if err := Watch(wg, cond); err != nil {
							println(err.Error())
							atomic.AddInt64(&fail, 1)
						} else {
							atomic.AddInt64(&success, 1)
						}
					}()
				}
				time.Sleep(time.Second)
				Emit(cond) // 触发
				wg.Wait()
				conns = nil
				println("success: ", success, "\tfail: ", fail)
			}
		}
	}
}
