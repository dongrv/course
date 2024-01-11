package filter

import (
	"context"
	"github.com/dongrv/toolkit"
	"regexp"
	"sync"
	"sync/atomic"
	"time"
)

type unique string

func (u unique) Bytes() []byte {
	return []byte(u)
}

// Clean 清洗数据
func (u unique) Clean() []byte {
	return []byte(regReplace(string(u)))
}

// Result 检测结果
type Result struct {
	Push  bool  // 是否推送报错
	Total int32 // 累计报错次数
}

func (r Result) Set(b bool, n int32) Result {
	r.Push = b
	r.Total = n
	return r
}

type metadata struct {
	raw     string // 原始消息
	unique  unique // 唯一标识(md5)
	create  int64  // 创建时间
	counter int32  // 计数器
}

// HashFilter 过滤器
type HashFilter struct {
	mu         sync.RWMutex
	store      map[unique]*metadata // 存储队列
	expiration int64                // 过期时间，举例：1800秒
	interval   float64              // 推送时间间隔，举例：300秒
	frequency  int32                // 推送频率
	clean      time.Ticker          // 定时器：清理数据
}

func NewHashFilter(expire int64, interval float64, freq int32) *HashFilter {
	return &HashFilter{
		store:      make(map[unique]*metadata, 100),
		expiration: expire,
		interval:   interval,
		frequency:  freq,
	}
}

// Do 执行检测过滤
func (h *HashFilter) Do(msg string) (r Result) {
	h.mu.Lock()
	defer h.mu.Unlock()

	hash := unique(toolkit.Md5(string(unique(msg).Clean())))
	now := time.Now().Unix()
	if m, ok := h.store[hash]; ok {
		if now-m.create >= h.expiration {
			// 复用过期键同时重置数据
			// 当前报错需推送
			m.raw, m.unique, m.create = msg, hash, time.Now().Unix()
			atomic.SwapInt32(&m.counter, 1)
			return r.Set(true, 1)
		}

		c := atomic.AddInt32(&m.counter, 1)
		if c%h.frequency == 0 {
			return r.Set(true, c) // 符合一定频率报错
		}
		return
	}
	h.store[hash] = &metadata{create: now, raw: msg, unique: hash, counter: 1} // 新建
	return r.Set(true, 1)
}

// Clean 清理过期元素
func (h *HashFilter) Clean() {
	h.mu.Lock()
	defer h.mu.Unlock()

	if len(h.store) == 0 {
		return
	}
	for u, m := range h.store {
		if time.Now().Unix()-m.create > h.expiration {
			m = nil
			delete(h.store, u)
		}
	}

}

// Tick 定时器：清理过期数据
func (h *HashFilter) Tick(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-h.clean.C:
			h.Clean()
		}
	}
}

// Deprecated:正则替换占位符，效率低
func regReplace(msg string) string {
	reg := regexp.MustCompile(`(%[(+|#)?\w)]{1,2})`)
	return string(reg.ReplaceAll([]byte(msg), []byte("")))
}
