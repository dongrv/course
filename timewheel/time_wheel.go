// Package timewheel
// 时间轮实现
package timewheel

import (
	"errors"
	"sync"
	"time"
)

const Unit = 100 * time.Millisecond // 单位

type Segment int64 // 时间片类型

type Type uint8 // 类型

const (
	_      uint8 = iota
	Single       // 单次任务
	Limit        // 需要执行多少次
	Loop         // 循环
)

type Wheel struct {
	Name string
	Type Type
	Pool []func()
}

// Tigger 触发任务
func (w *Wheel) Tigger() {
	for _, fn := range w.Pool {
		fn()
	}
}

type Timewheel struct {
	mu           sync.RWMutex
	Offset       int64               // 偏移量
	buckets      map[Segment][]Wheel // 轮子集合
	nameSeg      map[string]Segment  // 名称映射时间段
	BaseTimeline Segment             // 时间基线
}

func NewTimewheel() *Timewheel {
	return &Timewheel{
		buckets:      make(map[Segment][]Wheel, 1<<10),
		BaseTimeline: Segment(UnixMilli() / Unit),
	}
}

func (t *Timewheel) Add(w Wheel) error {
	seg := CurrentSeg()
	t.mu.Lock()
	defer t.mu.Unlock()
	if _, ok := t.buckets[seg]; !ok {
		t.buckets[seg] = make([]Wheel, 0, 1<<8)
	}
	if len(t.buckets[seg]) == cap(t.buckets[seg]) {
		return errors.New("current segment list in buckets is full")
	}
	t.buckets[seg] = append(t.buckets[seg], w)
	return nil
}

func (t *Timewheel) Remove(name string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	seg, ok := t.nameSeg[name]
	if !ok {
		return
	}
	for i, wheel := range t.buckets[seg] {
		if wheel.Name == name {
			buf := make([]Wheel, len(t.buckets[seg])-1)
			copy(buf[:i], t.buckets[seg][:i])
			copy(buf[i:], t.buckets[seg][i+1:])
			t.buckets[seg] = buf
			break
		}
	}
	delete(t.nameSeg, name)
}

// UnixMilli 当前毫秒时间戳
func UnixMilli() time.Duration {
	return time.Duration(time.Now().UnixMilli())
}

// CurrentSeg 当前 Segment
func CurrentSeg() Segment {
	return Segment(UnixMilli() / Unit)
}
