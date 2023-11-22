package testtcp

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

const Mode = 3

type Handler func(conn net.Conn, msg string, from string)

var container = map[int]Handler{
	1: handleConn1,
	2: handleConn2,
	3: handleConn3,
}

const addr = ":8086"

// HelloReq 请求，编号1
type HelloReq struct {
	Num int
	Msg string
}

// HelloResp 响应，编号2
type HelloResp struct {
	Payload string
}

// HelloReq2 请求，编号3
type HelloReq2 struct {
	SumList []int
}

// HelloResp2 响应，编号4
type HelloResp2 struct {
	Sum int
}

// 类型一：单向模式读、写
func rawRead(conn net.Conn, from string) error {
	for {
		buf := make([]byte, 256) // 单条消息限制字节大小
		n, err := conn.Read(buf)
		if err != nil {
			return err
		}
		send := &HelloReq{}
		_ = json.Unmarshal(buf[:n], send)
		println(fmt.Sprintf("%s received:%v\n", from, send))
	}
}

func rawWrite(conn net.Conn, from string) error {
	for {
		replay := &HelloResp{Payload: fmt.Sprintf("%s send:%d", from, time.Now().Second())}
		buf, err := json.Marshal(replay)
		if err != nil {
			return err
		}
		println(fmt.Sprintf("%s send:%v\n", from, replay))
		_, err = conn.Write(buf)
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
	}
}

// 类型一：处理连接读、写
func handleConn1(conn net.Conn, _, from string) {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	// 读数据
	go func() {
		defer wg.Done()
		println(rawRead(conn, from).Error())
	}()
	// 写数据
	go func() {
		defer wg.Done()
		println(rawWrite(conn, from).Error())
	}()
	wg.Wait()
}

// 类型二：同步请求+响应
func readProcess(conn net.Conn, _, from string) error {
	for {
		buf := make([]byte, 256) // 单条消息限制字节大小
		n, err := conn.Read(buf)
		if err != nil {
			return err
		}
		send := &HelloReq{}
		_ = json.Unmarshal(buf[:n], send)
		println(fmt.Sprintf("%s:received: %+v\n", from, send))
		// 响应
		send.Num++
		message, _ := json.Marshal(send)
		if err = writeProcess(conn, string(message), from); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
}

func writeProcess(conn net.Conn, msg, from string) error {
	if _, err := conn.Write([]byte(msg)); err != nil {
		return err
	}
	println(from + " write:" + msg + "\n")
	return nil
}

func handleConn2(conn net.Conn, msg, from string) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	// 读数据
	go func() {
		defer wg.Done()
		println(readProcess(conn, msg, from).Error())
	}()
	wg.Wait()
}

// 类型三：解析头字节获得消息长度

const (
	headerLen = 2    // 前两个字节表示当前总长度
	msgIdLen  = 1    // 消息编号
	bigEndian = true // 大端模式
)

func readConn(conn net.Conn, _ string, from string) {
	for {
		buf := make([]byte, headerLen)
		_, err := conn.Read(buf)
		if err != nil {
			Exit(err)
		}
		msgLen := ReadEndian(buf, headerLen, bigEndian)
		buf = make([]byte, msgLen)
		if _, err = io.ReadFull(conn, buf); err != nil {
			Exit(err)
		}
		msgSeq := ReadEndian(buf, msgIdLen, bigEndian) // 消息序号
		switch int(msgSeq) {
		case 1: // 解析请求消息1
			req := &HelloReq{}
			err := json.Unmarshal(buf[msgIdLen:], req)
			if err != nil {
				Exit(err)
			}
			fmt.Printf("%s:received->%#v\n", from, req)
			resp := HelloResp{Payload: "Hello:" + strconv.Itoa(rand.Intn(99))}
			respJson, err := json.Marshal(resp)
			if err = writeMsgConn(conn, 2, respJson, from); err != nil {
				Exit(err)
			}
		case 2: // 解析响应消息2
			req := &HelloResp{}
			err := json.Unmarshal(buf[msgIdLen:], req)
			if err != nil {
				Exit(err)
			}
			fmt.Printf("%s:received->%#v\n", from, req)
		case 3: // 解析请求消息3
			req := &HelloReq2{}
			err := json.Unmarshal(buf[msgIdLen:], req)
			if err != nil {
				Exit(err)
			}
			fmt.Printf("%s:received->%#v\n", from, req)
			sum := 0
			for _, v := range req.SumList {
				sum += v
			}
			resp := HelloResp2{Sum: sum}
			respJson, err := json.Marshal(resp)
			if err = writeMsgConn(conn, 4, respJson, from); err != nil {
				Exit(err)
			}
		case 4: // 解析响应消息4
			req := &HelloResp2{}
			err := json.Unmarshal(buf[msgIdLen:], req)
			if err != nil {
				Exit(err)
			}
			fmt.Printf("%s:received->%#v\n", from, req)
		}
	}

}

func writeMsgConn(conn net.Conn, msgId uint, msg []byte, _ string) error {
	newMsg := WrapHeader(WrapMsg(msg, msgId, msgIdLen), headerLen)
	if _, err := conn.Write(newMsg); err != nil {
		return err
	}
	return nil
}

func handleConn3(conn net.Conn, msg, from string) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		wg.Done()
		readConn(conn, "", from)
	}()
	wg.Wait()
}

func Serve() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	defer listener.Close()
	println("The server is running and listening at ", addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		go container[Mode](conn, "", "Server") // 选择类型
	}

}

func Exit(err error) {
	if err == nil {
		return
	}
	println(err.Error())
	os.Exit(-1)
}
