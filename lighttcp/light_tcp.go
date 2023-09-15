package lighttcp

import (
	"fmt"
	"io"
	"math"
	"net"
	"os"
	"sync"
	"time"
)

const (
	msgSeqLen = 2 // 消息编号长度，单位：字节
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
	headerLen int
	bigEndian bool // 是否大端
	MinLen    int  // 消息最小长度
	MaxLen    int  // 消息最大长度
}

// Init 设置消息体长度范围
func (mv *MsgProcessor) Init(bigEndian bool, head int) {
	if !(head == 1 || head == 2 || head == 4) {
		panic("args head must be 1,2,4")
	}
	mv.bigEndian, mv.headerLen, mv.MinLen = bigEndian, head, head
	mv.MaxLen = map[int]int{1: math.MaxInt8, 2: math.MaxInt16, 4: math.MaxInt32}[head]
}

// Validate 校验长度
func (mv *MsgProcessor) Validate(msgLen int) bool {
	return mv.MinLen <= msgLen && msgLen <= mv.MaxLen
}

// ReadEndian 读取数值
func (mv *MsgProcessor) ReadEndian(buf []byte) (msgLen uint) {
	return ReadEndian(buf, mv.headerLen, mv.bigEndian)
}

// WriteEndian 写入数值
func (mv *MsgProcessor) WriteEndian(msgLen uint, buf []byte) {
	WriteEndian(buf, mv.headerLen, msgLen, mv.bigEndian)
}

// 标准错误
func stdError(err error) {
	if err == nil {
		return
	}
	_, _ = fmt.Fprintf(os.Stderr, "%s\t报错信息：%s\n", time.Now().Format(time.DateTime+".000"), err.Error())
}

// 标准输出
func stdOut(b []byte) {
	_, _ = fmt.Fprintf(os.Stdout, "%s\t信息：%s\n", time.Now().Format(time.DateTime+".000"), b)
}

type Guard struct {
	wg          *sync.WaitGroup
	Processor   *MsgProcessor
	ReadCh      chan []byte // 读取管道
	WriteCh     chan []byte // 写入管道
	ActionQueue chan Action // 动作队列
	Collector   interface{} // TODO 数据收集器
}

func (g *Guard) Init() {
	g.Processor = &MsgProcessor{}
	g.Processor.Init(true, 4)
	g.wg = &sync.WaitGroup{}
}

func (g *Guard) Connect() {
	ln, err := net.Dial("tcp", ":2001")
	if err != nil {
		stdError(err)
	}
	g.wg.Add(2)
	go g.WriteTCP([]byte(`1234567890`), ln)
	go g.ReadTCP(ln)
	g.wg.Wait()
}

func (g *Guard) ReadTCP(conn net.Conn) {
	var err error

	defer func() {
		g.wg.Done()
		stdError(err)
	}()

	for {
		headerBuf := make([]byte, g.Processor.headerLen)
		_, err = io.ReadFull(conn, headerBuf)
		if err != nil {
			break
		}
		msgLen := g.Processor.ReadEndian(headerBuf)
		if msgLen == 0 {
			err = ErrReadEndianInvalid
			break
		}
		if !g.Processor.Validate(int(msgLen)) {
			err = fmt.Errorf("读取无效的消息长度：%d", msgLen)
			break
		}
		msgBuf := make([]byte, msgLen)
		_, err = conn.Read(msgBuf)
		if err != nil {
			break
		}
		stdOut(msgBuf)
		g.ReadCh <- msgBuf
	}
}

func (g *Guard) WriteTCP(msg []byte, conn net.Conn) {
	if len(msg) == 0 {
		return
	}
	var err error
	defer func() {
		g.wg.Done()
		stdError(err)
	}()

	for {
		t := g.Processor.headerLen + len(msg)
		if !g.Processor.Validate(t) {
			err = fmt.Errorf("读取无效的消息长度：%d", t)
			return
		}
		buf := make([]byte, t) // 消息总长度=头+消息体长度
		g.Processor.WriteEndian(uint(len(msg)), buf)
		copy(buf[g.Processor.headerLen:], msg)
		if _, err = conn.Write(buf); err != nil {
			break
		}
		stdOut(msg)
	}
}
