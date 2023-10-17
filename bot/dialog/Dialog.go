package dialog

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"net"
	"sync"
)

type IAction interface {
	Seq() int       // 动作序号
	Before()        // 执行前操作
	Do(f func(int)) // 执行
	After()         // 执行后操作
	IsLoop() bool   // 循环
	Message() proto.Message
}

type Event struct {
	SeqNum  int
	Msg     proto.Message
	LoopNum int // 循环次数
}

func (e *Event) Seq() int {
	return e.SeqNum
}

func (e *Event) Before()        {}
func (e *Event) Do(f func(int)) { f(e.SeqNum) }
func (e *Event) After()         {}
func (e *Event) IsLoop() bool {
	if e.LoopNum > 0 {
		e.LoopNum--
		return true
	}
	return false
}
func (e *Event) Message() proto.Message {
	return e.Msg
}

// 位置标记
const (
	OnSocket  At = `Start`
	OnRead    At = `OnRead`
	OnWrite   At = `OnWrite`
	OnProcess At = `OnProcess`
	OnSend    At = `OnSend`
)

// Processor 消息处理器
type Processor struct {
	headBytes int  // 消息长度数值存储的字节数
	bigEndian bool // 是否大端
	minLen    int  // 消息最小长度
	maxLen    int  // 消息最大长度
}

// Init 设置消息体长度范围
func (mv *Processor) Init() {
	mv.bigEndian = bigEndian
	mv.headBytes = headBytes
	mv.minLen = headBytes
	mv.maxLen = ByteMaxInt[headBytes]
}

type Dialog struct {
	wg        *sync.WaitGroup
	Processor *Processor
	ReadChan  chan []byte     // 读取管道
	WriteChan chan []byte     // 写入管道
	awaitChan chan IAction    // 动作队列
	MsgSync   *sync.WaitGroup // 消息同步
	//Collector   interface{} // TODO 数据收集器
}

func (d *Dialog) InputAction(i IAction) {
	d.awaitChan <- i
}

func NewDialog() *Dialog {
	r := &Dialog{
		wg:        new(sync.WaitGroup),
		Processor: &Processor{},
		ReadChan:  make(chan []byte, 10),
		WriteChan: make(chan []byte, 10),
		awaitChan: make(chan IAction, 10),
	}
	r.Processor.Init()
	return r
}

// Start 开始建立连接，启动Socket读写协程
func (d *Dialog) Start(ctx context.Context, ln net.Conn) {
	d.wg.Add(2)
	go d.OnWrite(ctx, ln)
	go d.OnRead(ctx, ln)
	d.Wait()
}

func (d *Dialog) Wait() {
	d.wg.Wait()
	logrus.Info("会话退出了")
}

func (d *Dialog) OnNotify(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			logrus.Debug("OnNotify goroutine exit")
			return
		case meta := <-d.awaitChan:
			meta.Before()
			d.OnSend(meta.Message(), uint(meta.Seq()))
			meta.After()
		}
	}
}

func (d *Dialog) OnProcess(ctx context.Context, closure func(msgId int32, msg []byte)) {
	for {
		select {
		case <-ctx.Done():
			logrus.Debug("OnProcess exit")
			return
		case msg := <-d.ReadChan:
			msgId := ReadEndian(msg, msgIdBytes, bigEndian) // 读取消息编号
			closure(int32(msgId), msg[msgIdBytes:])
		}
	}
}

func (d *Dialog) OnSend(pbMsg proto.Message, msgId uint) {
	msg, err := proto.Marshal(pbMsg)
	if err != nil {
		logrus.Errorf(`Dialog.OnSend err %s`, err.Error())
		return
	}
	d.WriteChan <- WrapHeader(WrapMsg(msg, msgId), headBytes)
}

func (d *Dialog) OnRead(ctx context.Context, conn net.Conn) {
	defer d.wg.Done()
	err := ReadSocket(ctx, d.ReadChan, conn, d.Processor.headBytes, func(i interface{}) bool {
		v, ok := i.(int)
		if !ok {
			return false
		}
		return d.Processor.minLen <= v && v <= d.Processor.maxLen
	})
	if err != nil {
		logrus.Errorf(`Dialog.OnRead err %s`, err.Error())
	}
}

func (d *Dialog) OnWrite(ctx context.Context, conn net.Conn) {
	defer d.wg.Done()
	if err := WriteSocket(ctx, d.WriteChan, conn); err != nil {
		logrus.Errorf(`Dialog.OnWrite err %s`, err.Error())
	}
}
