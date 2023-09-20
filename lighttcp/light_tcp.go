package lighttcp

import (
	"context"
	"net"
	"sync"
)

type Action interface {
	SEQ() int    // 动作序号
	Before()     // 执行前操作
	Do()         // 执行
	After()      // 执行后操作
	Cycle() bool // 循环？
}

// MsgProcessor 消息处理器
type MsgProcessor struct {
	msgHeadSize int  // 消息长度数值存储的字节数
	bigEndian   bool // 是否大端
	minLen      int  // 消息最小长度
	maxLen      int  // 消息最大长度
}

// 位置标记
const (
	OnSocket At = `OnSocket`
	OnRead   At = `OnRead`
	OnWrite  At = `OnWrite`
)

// Init 设置消息体长度范围
func (mv *MsgProcessor) Init() {
	mv.bigEndian = bigEndian
	mv.msgHeadSize = msgHeaderSize
	mv.minLen = msgHeaderSize
	mv.maxLen = ByteMaxInt[msgHeaderSize]
}

type Guard struct {
	wg          *sync.WaitGroup
	Processor   *MsgProcessor
	ReadCh      chan []byte // 读取管道
	WriteCh     chan []byte // 写入管道
	ActionQueue chan Action // 动作队列
	//Collector   interface{} // TODO 数据收集器
}

func (g *Guard) Init() {
	g.Processor = &MsgProcessor{}
	g.Processor.Init()
	g.wg = &sync.WaitGroup{}
	g.ReadCh = make(chan []byte, 1)
	g.WriteCh = make(chan []byte, 1)
	g.ActionQueue = make(chan Action, 100)
}

// TCPConfig 通讯配置
type TCPConfig struct {
	Network string
	Addr    string
}

// OnSocket 建立连接，启动Socket读写协程
func (g *Guard) OnSocket(ctx context.Context, c TCPConfig) {
	ln, err := net.Dial(c.Network, c.Addr)
	if err != nil {
		stdError(err, OnSocket)
		return
	}
	defer ln.Close()

	g.wg.Add(2)
	go g.OnWrite(ctx, ln)
	go g.OnRead(ctx, ln)
	g.wg.Wait()
}

func (g *Guard) Process() {
	for {
		select {
		case _ = <-g.ReadCh:

		}
	}
}

func (g *Guard) OnRead(ctx context.Context, conn net.Conn) {
	defer g.wg.Done()
	stdError(ReadSocket(ctx, g.ReadCh, conn, g.Processor.msgHeadSize, func(i interface{}) bool {
		v, ok := i.(int)
		if !ok {
			return false
		}
		return g.Processor.minLen <= v && v <= g.Processor.maxLen
	}), OnRead)
}

func (g *Guard) OnWrite(ctx context.Context, conn net.Conn) {
	defer g.wg.Done()
	stdError(WriteSocket(ctx, g.WriteCh, conn), OnWrite)
}
