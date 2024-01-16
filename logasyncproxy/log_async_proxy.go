package logasyncproxy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

type State uint8 // 状态标记

const (
	Closed State = iota
	Running
)

// Meta 日志参数单元
type Meta struct {
	Level  int32         // 级别
	Format string        // 文本
	Args   []interface{} // 参数
}

type Couriers interface {
	Read() Meta
}

type Publisher interface {
	Publish()
}

const (
	bufferLen           = 100                    // 缓冲区刷新计数
	chanCap             = 1 << 10                // 管道长度
	flushBufferInterval = 100 * time.Millisecond // 刷新间隔
	heartbeat           = 10 * time.Millisecond
)

type AsyncProxy struct {
	wg            sync.WaitGroup
	Writer        io.Writer
	state         State         // 运行状态
	queue         chan string   // 传输管道
	bufferLock    sync.Mutex    // 缓冲区锁
	buffer        bytes.Buffer  // 缓冲区
	bufferLen     int           // 写入缓冲区消息个数
	flushTime     int64         // 最近刷新缓存时间：毫秒
	Heartbeat     *time.Ticker  // 心跳检测
	FlushInterval time.Duration // 刷新间隔时间
	stop          chan struct{} // 停止信号
}

func NewAsyncProxy(w io.Writer) *AsyncProxy {
	return &AsyncProxy{
		Writer:        w,
		queue:         make(chan string, chanCap),
		buffer:        bytes.Buffer{},
		flushTime:     time.Now().UnixNano() / int64(time.Millisecond),
		Heartbeat:     time.NewTicker(heartbeat),
		FlushInterval: flushBufferInterval,
		stop:          make(chan struct{}, 1),
	}
}

// destruct 析构资源
func (ap *AsyncProxy) destruct() {
	ap.Heartbeat.Stop()
}

func (ap *AsyncProxy) write(s string) {
	ap.bufferLock.Lock()
	ap.buffer.WriteString(s)
	ap.bufferLen++
	ap.bufferLock.Unlock()
	if ap.bufferLen == bufferLen {
		ap.Flush()
	}
}

// Flush 刷新缓存数据到文件
func (ap *AsyncProxy) Flush() {
	ap.bufferLock.Lock()
	defer ap.bufferLock.Unlock()
	if ap.bufferLen == 0 {
		return
	}
	if _, err := ap.buffer.WriteTo(ap.Writer); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "log async proxy error:%s", err.Error())
	}
	ap.bufferLen = 0
	atomic.SwapInt64(&ap.flushTime, time.Now().UnixNano()/1e6)
}

// Publish 发布日志
func (ap *AsyncProxy) Publish(meta Meta) {
	if ap.state == Closed {
		return
	}
	w := wrap{
		Level: meta.Level, Time: time.Now().Format(time.RFC3339),
		Msg: fmt.Sprintf(meta.Format, meta.Args...),
	}
	if len(ap.queue) >= chanCap {
		_, _ = fmt.Fprintln(os.Stdout, "log plugin AsyncProxy channel is full")
		return
	}
	ap.queue <- LineEnd(w.String())
}

// Run 接收和处理日志
func (ap *AsyncProxy) Run(ctx context.Context) {
	ap.wg.Add(1)
	defer func() {
		ap.wg.Done()
	}()
	for {
		select {
		case s := <-ap.queue:
			ap.write(s)
		case <-ap.Heartbeat.C:
			if time.Now().UnixNano()/1e6-atomic.LoadInt64(&ap.flushTime) >= int64(flushBufferInterval)/1e6 {
				ap.Flush()
			}
		case <-ap.stop:
			ap.emptyQueue(Closed)
			return
		case <-ctx.Done():
			ap.emptyQueue(Closed)
			return
		}
	}
}

func (ap *AsyncProxy) Wait() {
	ap.wg.Wait()
}

// 清空队列元素
func (ap *AsyncProxy) emptyQueue(state State) {
	if ap.state != Running {
		return
	}
	ap.state = state
	if len(ap.queue) > 0 {
		for q := range ap.queue {
			ap.write(q)
		}
	}
	if ap.bufferLen > 0 {
		ap.Flush()
	}
}

func (ap *AsyncProxy) Running() bool {
	return ap.state == Running
}

func (ap *AsyncProxy) Stop() {
	ap.stop <- struct{}{}
	close(ap.stop)
	ap.Heartbeat.Stop()
}

type Manager struct {
	mu   sync.RWMutex
	Pool map[int32]*AsyncProxy
}

func (m *Manager) Publish(meta Meta) {
	m.mu.RLock()
	p, ok := m.Pool[meta.Level]
	m.mu.RUnlock()
	if !ok {
		// TODO 注册新文件
	}
	p.Publish(meta)
}

type wrap struct {
	Level int32  `json:"level"`
	Msg   string `json:"msg"`
	Time  string `json:"time"`
}

func (w *wrap) String() string {
	bs, _ := json.Marshal(w)
	return string(bs)
}

func LineEnd(str string) string {
	return str + "\n"
}
