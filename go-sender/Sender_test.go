package v2

import "testing"

var _ Processor = (*example)(nil)

type example struct{ *Base }

var topic = Topic("example")

func (e *example) Topic() Topic        { return topic }
func (e *example) String() string      { return "I'm an example." }
func (e *example) Description() string { return "" }

func TestPool_Send(t *testing.T) {
	Get().Register(&Option{
		Topic:        topic,
		GoroutineNum: 10,
		ChannelCap:   1 << 10,
		Call:         NewCall(1 << 10).SetSingle(func(processor Processor) { println(processor.String()) }),
	})

	for i := 0; i < 100; i++ {
		Get().Send(&example{})
	}
}
