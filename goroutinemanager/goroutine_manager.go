package goroutinemanager

import (
	"context"
)

type Runner interface {
	Run(ctx context.Context)
	Stop()
	Wait()
}

var defaultGoroutine = &Goroutine{rs: make(map[string][]Runner)}

type Goroutine struct {
	rs map[string][]Runner
}

func (g *Goroutine) Push(key string, r Runner) *Goroutine {
	if _, ok := g.rs[key]; !ok {
		g.rs[key] = make([]Runner, 0, 2)
	}
	g.rs[key] = append(g.rs[key], r)
	return g
}

// Run 拉起协程
func (g *Goroutine) Run(key string, ctx context.Context) {
	if _, ok := g.rs[key]; !ok {
		return
	}
	for _, r := range g.rs[key] {
		go r.Run(ctx)
	}
}

// Stop 停止协程
func (g *Goroutine) Stop(key string) {
	if key == "" {
		return
	}
	if _, ok := g.rs[key]; !ok {
		return
	}
	for _, r := range g.rs[key] {
		r.Stop()
		r.Wait()
	}
}

// StopAll 停止所有协程
func (g *Goroutine) StopAll() {
	if g.rs == nil {
		return
	}
	for k := range g.rs {
		g.Stop(k)
	}
}

// Get 获取协程管理器
func Get() *Goroutine {
	return defaultGoroutine
}
