package binarytest

import (
	"aaagame/tests/course/binarytest/protocol"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net"
	"sync"
	"time"
	"unsafe"
)

/*
大端和小端

什么是字节序？
字节（Byte）是计算机世界的计量单位，比如，一部电视剧是10G个字节（1GB）,一张图片是1K个字节（1KB）。这些数据量有多有少，大多数时候，一个字节肯定是装不下的，这个时候我们就要用到多字节。在使用多字节的存的时候，就会涉及到一个顺序问题。
在计算机中，字节序是指多字节数据在计算机内存中存储或者网络传输时各字节的存储顺序。有两种储存数据的方式：大端字节序（big endian）和小端字节序（little endian）。举例来说，数值0x2211使用两个字节储存：高位字节是0x22，低位字节是0x11。

什么是大端序和小端序？
大端字节序：高位字节在前，低位字节在后，这是人类读写数值的方法。简单来说，就是按照从低地址到高地址的顺序存放数据的高位字节到低位字节，就如同例子中的0X2211
小端字节序：低位字节在前，高位字节在后，就是按照从低地址到高地址的顺序存放据的低位字节到高位字节。即以0x1122形式储存。
*/

const (
	StreamHeadSize = 4 // 消息头字节长度
)

func Binary() {
	var i uint32 = 380
	iptr := (*[4]byte)(unsafe.Pointer(&i))
	for _, ptr := range iptr {
		fmt.Printf("%02X", ptr)
	}
	fmt.Println()

	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, i)
	for _, bin := range b {
		fmt.Printf("%02X", bin)
	}
	fmt.Println()

	binary.BigEndian.PutUint32(b, i)
	for _, bin := range b {
		fmt.Printf("%02X", bin)
	}
	fmt.Println()
}

type endianMode int

const (
	big endianMode = iota
	little
)

func TCPServer() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	for {
		ln, netErr := l.Accept()
		if netErr != nil {
			if err, ok := netErr.(net.Error); ok && err.Timeout() {
				// TODO log
				continue
			}
			fmt.Errorf("create link error:%s", err.Error())
			continue
		}
		fmt.Printf("server is ruinning at %s\n", addr)
		go func(conn net.Conn) {
			var wg sync.WaitGroup
			wg.Add(2)
			msgChan := make(chan *protocol.SayResp, 1)
			// read
			go func() {
				defer wg.Done()
				for {
					buf := make([]byte, StreamHeadSize)
					_, err := io.ReadFull(conn, buf)
					if err != nil {
						fmt.Printf("read error:%s\n", err.Error())
						return
					}
					msgLen := binary.LittleEndian.Uint32(buf[:]) // 实际消息长度
					streamBuf := make([]byte, msgLen)
					n, err := conn.Read(streamBuf)
					if err != nil {
						fmt.Printf("服务器读取消息错误:%s\n", err.Error())
						return
					}
					req := &protocol.SayReq{}
					err = json.Unmarshal(streamBuf[:n], req)
					if err != nil {
						fmt.Printf("服务器读取消息报错:%s\n", err.Error())
						return
					}
					fmt.Printf("服务器读取消息内容:%v\n", req)
					resp := protocol.SayResp{}
					if req.Id == 0 {
						resp = protocol.SayResp{
							Err:     0,
							Message: "FromId is 0",
							Payload: &protocol.SayResp_Payload{
								From: 0,
								Msg:  "I'm payload message",
							},
						}
					} else {
						resp = protocol.SayResp{
							Err:     1,
							Message: "FromId is 0",
							Payload: &protocol.SayResp_Payload{
								From: 1,
								Msg:  "I'm payload message1",
							},
						}
					}
					msgChan <- &resp
				}
			}()
			// write
			go func() {
				ticker := time.NewTicker(2 * time.Second)
				defer func() {
					wg.Done()
					ticker.Stop()
				}()

				for {
					select {
					case resp := <-msgChan:
						msg, err := json.Marshal(resp)
						if err != nil {
							fmt.Printf("服务器写入错误:%s\n", err.Error())
							return
						}
						msgLen := len(msg)                                                  // 获取消息长度
						buf := make([]byte, StreamHeadSize+msgLen)                          // 构造缓冲区，长度=头长度+消息体长度
						binary.LittleEndian.PutUint32(buf[:StreamHeadSize], uint32(msgLen)) // 消息体长度写入头部字节
						copy(buf[StreamHeadSize:], msg)                                     // 实际消息体copy到缓冲区
						if _, err := conn.Write(buf); err != nil {
							fmt.Printf("服务器写入错误:%s\n", err.Error())
							return
						}
						//fmt.Printf("服务器写入心跳计数:%d\n", tickNum)
					}
				}
			}()
			wg.Wait()
			fmt.Println("goroutine exit.")
		}(ln)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Client() {
	time.Sleep(1 * time.Second)
	addr := ":8080"
	ln, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("dial success,host:%s\n", addr)
	var ch = make(chan struct{}, 0)
	go func(conn net.Conn) {
		var wg sync.WaitGroup
		wg.Add(2)
		// read
		go func() {
			defer wg.Done()
			for {
				buf := make([]byte, StreamHeadSize)
				_, err := io.ReadFull(conn, buf)
				if err != nil {
					fmt.Printf("客户端读取错误: %s\n", err.Error())
					return
				}
				stream := make([]byte, binary.LittleEndian.Uint32(buf[:]))
				_, err = conn.Read(stream)
				if err != nil {
					fmt.Printf("客户端读取错误: %s\n", err.Error())
					return
				}
				resp := &protocol.SayResp{}
				err = json.Unmarshal(stream, resp)
				if err != nil {
					fmt.Printf("客户端读取报错:%s", err.Error())
					return
				}
				fmt.Printf("客户端读取消息内容:%v\n", resp)
				time.Sleep(time.Second)
			}
		}()
		// write
		go func() {
			defer wg.Done()
			for {
				req := &protocol.SayReq{
					Message: "我是来自客户端的推送信息",
					Id:      int32(rand.Intn(2)),
				}
				msg, err := json.Marshal(req)
				if err != nil {
					fmt.Printf("客户端写入错误：%s\n", err.Error())
					return
				}
				msgLen := len(msg)
				buf := make([]byte, StreamHeadSize+msgLen)
				binary.LittleEndian.PutUint32(buf[:StreamHeadSize], uint32(msgLen))
				copy(buf[StreamHeadSize:], msg)
				_, err = ln.Write(buf)
				if err != nil {
					fmt.Printf("客户端写入错误:%s\n", err.Error())
					return
				}
				time.Sleep(5 * time.Second)
			}
		}()
		wg.Wait()
		ch <- struct{}{}
	}(ln)
	<-ch
}
