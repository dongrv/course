package lighttcp

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"os"
	"time"
)

var (
	ErrReadEndianInvalid = errors.New("无效的端读取操作")
	ErrValidateFailed    = errors.New("校验参数失败")
)

const (
	printStdLog = true // 是否输出标准日志
	bigEndian   = true // 大端模式

	msgHeaderSize = 1 << 2 // 消息头占用字节
	msgIdUsedSize = 1 << 1 // 消息号占用字节

)

// ByteMaxInt 字节对应最大值
var ByteMaxInt = map[int]int{
	1: math.MaxInt8,
	2: math.MaxInt16,
	4: math.MaxInt32,
	8: math.MaxInt64,
}

type At string

func (a At) String() string {
	return string(a)
}

// StdError 标准错误
func StdError(err error, at At) {
	if !printStdLog || err == nil {
		return
	}
	_, _ = fmt.Fprintf(os.Stderr, "%s\t%s\t错误：%s\n", time.Now().Format(time.DateTime+".000"), at.String(), err.Error())
}

// StdOut 标准输出
func StdOut(b []byte, at At) {
	if !printStdLog {
		return
	}
	_, _ = fmt.Fprintf(os.Stdout, "%s\t%s\t输出：%s\n", time.Now().Format(time.DateTime+".000"), at, b)
}

// ReadEndian 读取消息长度
func ReadEndian(buf []byte, headerLen int, bigEndian bool) (msgLen uint) {
	if len(buf) < headerLen {
		return 0 // 消息长度不够
	}
	switch headerLen {
	case 1:
		msgLen = uint(buf[0])
	case 2:
		if bigEndian {
			msgLen = uint(binary.BigEndian.Uint16(buf[:2]))
		} else {
			msgLen = uint(binary.LittleEndian.Uint16(buf[:2]))
		}
	case 4:
		if bigEndian {
			msgLen = uint(binary.BigEndian.Uint32(buf[:4]))
		} else {
			msgLen = uint(binary.LittleEndian.Uint32(buf[:4]))
		}
	case 8:
		if bigEndian {
			msgLen = uint(binary.BigEndian.Uint64(buf[:8]))
		} else {
			msgLen = uint(binary.LittleEndian.Uint64(buf[:8]))
		}
	}
	return msgLen
}

// WriteEndian 消息长度写入头部字节
func WriteEndian(buf []byte, byteLen int, msgLen uint, bigEndian bool) {
	switch byteLen {
	case 1:
		buf[0] = byte(msgLen)
	case 2:
		if bigEndian {
			binary.BigEndian.PutUint16(buf[:2], uint16(msgLen))
		} else {
			binary.LittleEndian.PutUint16(buf[:2], uint16(msgLen))
		}
	case 4:
		if bigEndian {
			binary.BigEndian.PutUint32(buf[:4], uint32(msgLen))
		} else {
			binary.LittleEndian.PutUint32(buf[:4], uint32(msgLen))
		}
	case 8:
		if bigEndian {
			binary.BigEndian.PutUint64(buf[:8], uint64(msgLen))
		} else {
			binary.LittleEndian.PutUint64(buf[:8], uint64(msgLen))
		}
	}
}

type ValidateFunc func(i interface{}) bool // 校验函数模板

func ReadSocket(ctx context.Context, ch chan<- []byte, conn net.Conn, bytes int, fs ...ValidateFunc) error {
	var err error
	for {
		select {
		case <-ctx.Done():
			err = nil
			break
		default:
		}
		headBuf := make([]byte, bytes)
		_, err = io.ReadFull(conn, headBuf)
		if err != nil {
			break
		}
		msgLen := ReadEndian(headBuf, bytes, bigEndian)
		if msgLen == 0 {
			err = ErrReadEndianInvalid
			break
		}
		for _, f := range fs {
			if !f(msgLen) {
				return ErrValidateFailed
			}
		}
		buf := make([]byte, msgLen)
		_, err = conn.Read(buf)
		if err != nil {
			break
		}
		StdOut(buf, `ReadSocket`)
		ch <- buf
	}
	return err
}

func WriteSocket(ctx context.Context, ch <-chan []byte, conn net.Conn) error {
	for {
		select {
		case msg := <-ch:
			if len(msg) == 0 {
				continue
			}
			if _, err := conn.Write(msg); err != nil {
				return err
			}
			StdOut(msg, `WriteSocket`)
		case <-ctx.Done():
			StdOut([]byte("上文环境退出"), `WriteSocket`)
			return nil
		}
	}
}

// WrapHeader 包装消息头
func WrapHeader(msg []byte, headLen int) []byte {
	buf := make([]byte, len(msg)+headLen)
	WriteEndian(buf, headLen, uint(len(msg)), bigEndian)
	copy(buf[headLen:], msg)
	return buf
}

// WrapMsg 包装消息和消息编号
func WrapMsg(msg []byte, msgId uint) []byte {
	buf := make([]byte, len(msg)+msgIdUsedSize)
	WriteEndian(buf, msgIdUsedSize, msgId, bigEndian)
	copy(buf[2:], msg)
	return buf
}
