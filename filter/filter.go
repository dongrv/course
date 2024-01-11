package filter

import (
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
func (u unique) Clean(replace string) []byte {
	reg := regexp.MustCompile(replace)
	return reg.ReplaceAll(u.Bytes(), []byte(""))
}

type metadata struct {
	raw     string    // 原始消息
	unique  unique    // 唯一标识
	create  time.Time // 创建时间
	counter int32     // 计数器
}

// HashFilter 过滤器
type HashFilter struct {
	mu         sync.RWMutex
	queue      map[unique]*metadata // 存储队列
	expiration float64              // 过期时间，举例：1800秒
	frequency  int32                // 推送频率
	interval   int64                // 推送时间间隔，举例：300秒
}

// Decision 检测结果
type Decision struct {
	Push  bool  // 是否推送报错
	Total int32 // 累计报错次数
}

func (d Decision) Set(b bool, n int32) Decision {
	d.Push = b
	d.Total = n
	return d
}

// Do 检测
func (f *HashFilter) Do(msg string) (d Decision) {
	f.mu.Lock()
	defer f.mu.Unlock()

	hash := unique(toolkit.Md5(string(unique(msg).Clean("(%[(+|#)?\\w)]{1,2})"))))

	if m, ok := f.queue[hash]; ok {
		diff := time.Since(m.create).Seconds()
		if diff >= f.expiration {
			// 复用过期键同时重置数据
			// 当前报错需推送
			m.raw = msg
			m.unique = hash
			m.create = time.Now()
			atomic.SwapInt32(&m.counter, 1)
			return d.Set(true, 1)
		}

		c := atomic.AddInt32(&m.counter, 1)
		if c%f.frequency == 0 {
			return d.Set(true, c) // 符合一定频率报错
		}
		return
	}
	f.queue[hash] = &metadata{create: time.Now(), raw: msg, unique: hash, counter: 1} // 新建
	return d.Set(true, 1)
}
