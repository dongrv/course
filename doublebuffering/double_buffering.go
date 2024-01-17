package doublebuffering

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// 双缓存实现原型逻辑

const BufferCap = 10 // 容量

const (
	write int64 = iota
	read
)

type Buffer struct {
	ch    chan byte
	state atomic.Int64
}

type Exchange struct {
	left, right *Buffer // 写、读
	lock        atomic.Uint64
	wg          sync.WaitGroup
}

func NewExchange() *Exchange {
	return &Exchange{
		left:  &Buffer{ch: make(chan byte, BufferCap)},
		right: &Buffer{ch: make(chan byte, BufferCap)},
	}
}

func (e *Exchange) block(delta int) {
	e.wg.Add(delta)
	e.wg.Wait()
}

func (e *Exchange) release() {
	e.wg.Done()
}

func (e *Exchange) Write(b byte) bool {
	if e.left.state.Load() != write && e.right.state.Load() != write {
		e.lock.Store(1)
		e.block(1)
		e.lock.CompareAndSwap(1, 0)
	}

	fmt.Println("写：", b)

	if e.left.state.Load() == write {
		e.left.ch <- b
		if len(e.left.ch) == BufferCap {
			fmt.Println("left满了")
			e.left.state.CompareAndSwap(write, read)
			fmt.Println("left wait to read")
		}
		return true
	}
	if e.right.state.Load() == write {
		e.right.ch <- b
		if len(e.right.ch) == BufferCap {
			fmt.Println("right满了")
			e.right.state.CompareAndSwap(write, read)
			fmt.Println("right wait to read")
		}
		return true
	}

	return true
}

func (e *Exchange) Read() bool {
	if e.left.state.Load() == read {
		for {
			select {
			case v := <-e.left.ch:
				fmt.Println("读left:", v)
				time.Sleep(1 * time.Second)
				if len(e.left.ch) == 0 {
					e.left.state.CompareAndSwap(read, write)
					if e.lock.Load() == 1 {
						e.release()
					}
					fmt.Println("left switch state to read")
					return true
				}
			}
		}
	}
	if e.right.state.Load() == read {
		for {
			select {
			case v := <-e.right.ch:
				fmt.Println("读right:", v)
				time.Sleep(1 * time.Second)
				if len(e.right.ch) == 0 {
					e.right.state.CompareAndSwap(read, write)
					fmt.Println("right switch state to read")
					if e.lock.Load() == 1 {
						e.wg.Done()
					}
					return true
				}
			}
		}
	}
	return true
}

func (e *Exchange) RWGoroutine() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	t := time.Duration(10)
	go func() {
		defer wg.Done()
		for e.Write(byte(rand.Intn(256))) {
			time.Sleep(t * time.Millisecond)
		}
	}()
	go func() {
		defer wg.Done()
		for e.Read() {
			time.Sleep(t * time.Millisecond)
		}
	}()
	wg.Wait()
}
