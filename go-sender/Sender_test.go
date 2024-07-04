package v2

import (
	"sync/atomic"
	"testing"
	"time"
)

var _ Processor = (*example)(nil)

type example struct{ *Base }

var topic = Topic("example")

func (e *example) Topic() Topic        { return topic }
func (e *example) String() string      { return "I'm an example." }
func (e *example) Description() string { return "" }

var num int64

func TestPool_Send(t *testing.T) {
	Get().Register(&Option{
		Topic:        topic,
		GoroutineNum: 10,
		ChannelCap:   1 << 10,
		Call: NewCall(1 << 10).SetSingle(func(processor Processor) {
			atomic.AddInt64(&num, 1)
		}),
	})

	for i := 0; i < 10000; i++ {
		Get().Send(&example{})
	}

	time.Sleep(10 * time.Second)
	println(num)
}
