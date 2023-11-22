package testtcp

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

const Mode = 2

type Handler func(conn net.Conn, msg string, from string)

var container = map[int]Handler{
	1: handleRawConn,
	2: handleProcessConn,
}

const addr = ":8086"

// Send 发送消息
type Send struct {
	Num int
	Msg string
}

// Replay 恢复消息
type Replay struct {
	Msg string
}

// 类型一：单向模式读、写
func rawRead(conn net.Conn, from string) error {
	for {
		buf := make([]byte, 256) // 单条消息限制字节大小
		n, err := conn.Read(buf)
		if err != nil {
			return err
		}
		send := &Send{}
		_ = json.Unmarshal(buf[:n], send)
		println(fmt.Sprintf("%s received:%v\n", from, send))
	}
}

func rawWrite(conn net.Conn, from string) error {
	for {
		replay := &Replay{Msg: fmt.Sprintf("%s send:%d", from, time.Now().Second())}
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
func handleRawConn(conn net.Conn, _, from string) {
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
		send := &Send{}
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

func handleProcessConn(conn net.Conn, msg, from string) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	// 读数据
	go func() {
		defer wg.Done()
		println(readProcess(conn, msg, from).Error())
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
