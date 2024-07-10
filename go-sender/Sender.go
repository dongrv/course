package v2

import (
	"fmt"
	"github.com/dongrv/iterator"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Requester interface {
	Topic() Topic
	Header() http.Header
	Do() error
	Backup()
	Raw() proto.Message
}

type StringReader interface {
	String() string
}

type BytesReader interface {
	Bytes() []byte
}

type Processor interface {
	Requester
	StringReader
	BytesReader
}

type Base struct{} // 默认实现

func (b *Base) Topic() Topic        { return "" }
func (b *Base) Header() http.Header { return nil }
func (b *Base) Do() error           { return nil }
func (b *Base) Backup()             {}
func (b *Base) Raw() proto.Message  { return nil }
func (b *Base) String() string      { return "" }
func (b *Base) Bytes() []byte       { return nil }

type ChanProcessor interface {
	Size() int
	Idle() bool
	Push(Processor) bool
	Pop() <-chan Processor
	Close()
}

type Channel struct {
	ch chan Processor
}

func NewChannel(cap int) *Channel {
	return &Channel{ch: make(chan Processor, cap)}
}

func (c *Channel) Size() int  { return len(c.ch) }
func (c *Channel) Close()     { close(c.ch) }
func (c *Channel) Idle() bool { return len(c.ch) < cap(c.ch) }
func (c *Channel) Push(p Processor) bool {
	if !c.Idle() {
		return false
	}
	c.ch <- p
	return true
}
func (c *Channel) Pop() <-chan Processor { return c.ch }

type coroutine struct {
	channel ChanProcessor
	stop    chan struct{}
	live    time.Time
	call    *Call
}

type SingleFunc func(processor Processor)
type BatchFunc func(processors []Processor)

type Handler interface {
	Handle(processor Processor)
}

type Call struct {
	single SingleFunc
	batch  BatchFunc
	value  []Processor
	cap    int
}

func NewCall(cap int) *Call { return &Call{value: make([]Processor, 0, cap), cap: cap} }

func (c *Call) SetSingle(s SingleFunc) *Call {
	c.single = s
	return c
}

func (c *Call) SetBatch(b BatchFunc) *Call {
	c.batch = b
	return c
}

func (c *Call) clear() {
	if c.batch != nil && len(c.value) > 0 {
		c.batch(c.value)
	}
}

func (c *Call) Handle(processor Processor) {
	if c.single != nil {
		c.single(processor)
		return
	}
	if c.batch == nil {
		return
	}
	c.value = append(c.value, processor)
	if len(c.value) >= c.cap {
		c.batch(c.value)
		c.value = c.value[:0]
	}
}

func newCoroutine(wg *sync.WaitGroup, op *Option) *coroutine {
	g := &coroutine{
		channel: NewChannel(op.ChannelCap),
		stop:    make(chan struct{}),
		call:    op.Call,
	}
	go g.Run(wg)
	return g
}

func (g *coroutine) Run(wg *sync.WaitGroup) {
	wg.Add(1)
	//println("启动了一个协程")
	defer wg.Done()
	timer := time.NewTimer(time.Duration(1500+rand.Intn(1500)) * time.Millisecond) // 增加随机性
	defer timer.Stop()
	for {
		select {
		case meta := <-g.channel.Pop():
			g.call.Handle(meta)
			g.live = time.Now()
		case <-timer.C:
			g.clear()
		case <-g.stop:
			g.clear()
			return
		}
	}
}

func (g *coroutine) clear() {
	if g.channel.Size() > 0 {
		for meta := range g.channel.Pop() { // 清空管道
			g.call.Handle(meta)
			if g.channel.Size() == 0 {
				break
			}
		}
	}
	g.call.clear()
}

func (g *coroutine) Quit() {
	g.stop <- struct{}{}
	close(g.stop)
	g.channel.Close()
}

type group struct {
	wg   *sync.WaitGroup
	list []*coroutine
	iter iterator.Func
}

func newGroup(op *Option) *group {
	g := &group{
		wg:   &sync.WaitGroup{},
		list: make([]*coroutine, op.GoroutineNum),
		iter: iterator.Get(),
	}
	for i := range g.list {
		g.list[i] = newCoroutine(g.wg, op)
	}
	return g
}

// Dispatch 遍历分发消息
func (g *group) Dispatch(p Processor) bool {
	return (g.list[int(g.iter())%len(g.list)]).channel.Push(p)
}

func (g *group) Quit() {
	if len(g.list) == 0 {
		return
	}
	for _, goroutine := range g.list {
		goroutine.Quit()
	}
	g.wg.Wait()
}

type Topic string // 主题

func (t Topic) String() string {
	return string(t)
}

type Pool struct {
	mu    sync.RWMutex
	sort  []Topic
	store map[Topic]*group
	state bool
}

func New() *Pool {
	return &Pool{sort: make([]Topic, 0), store: map[Topic]*group{}, state: true}
}

var defaultPool *Pool

func Get() *Pool {
	if defaultPool != nil {
		return defaultPool
	}
	defaultPool = New()
	return defaultPool
}

type Option struct {
	Topic        Topic
	GoroutineNum int
	ChannelCap   int
	Call         *Call
}

func (pool *Pool) Register(op *Option) {
	pool.mu.Lock()
	defer pool.mu.Unlock()
	if _, ok := pool.store[op.Topic]; ok {
		return
	}
	pool.sort = append(pool.sort, op.Topic)
	pool.store[op.Topic] = newGroup(op)
}

func (pool *Pool) Send(p Processor) bool {
	if !pool.state {
		return false
	}
	pool.mu.RLock()
	defer pool.mu.RUnlock()
	g, ok := pool.store[p.Topic()]
	if !ok {
		return false
	}
	return g.Dispatch(p)
}

func (pool *Pool) Quit() {
	pool.state = false
	if len(pool.store) == 0 {
		return
	}
	for _, topic := range pool.sort {
		pool.store[topic].Quit() // 按照注册顺序退出
		fmt.Printf("The topic:%s exit\n", topic)
	}
	fmt.Print("All topics exit\n")
}
