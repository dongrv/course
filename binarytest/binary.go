package binarytest

import (
	"encoding/binary"
	"fmt"
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
	addr := ":8080"
	l, err := net.Listen("tcp", addr)
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
			// read
			go func() {
				defer wg.Done()
				for {
					buf := make([]byte, 4)
					_, err := conn.Read(buf)
					if err != nil {
						fmt.Printf("read error:%s\n", err.Error())
						return
					}
					v := binary.LittleEndian.Uint32(buf[:4])
					fmt.Printf("server receive from client:%d\n", v)
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
					case <-ticker.C:
						ticker.Reset(2 * time.Second)
						buf := make([]byte, 4)
						var v uint32 = 100
						binary.LittleEndian.PutUint32(buf[:4], v)
						if _, err := conn.Write(buf); err != nil {
							fmt.Printf("write error:%s\n", err.Error())
							return
						}
						fmt.Printf("server write %v\n", v)
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
				buf := make([]byte, 4)
				_, err := ln.Read(buf)
				if err != nil {
					fmt.Printf("client read error: %s\n", err.Error())
					return
				}
				v := binary.LittleEndian.Uint32(buf[:])
				fmt.Printf("client read:%d\n", v)
			}
		}()
		// write
		go func() {
			defer wg.Done()
			for {
				buf := make([]byte, 4)
				var v = rand.Uint32()
				binary.LittleEndian.PutUint32(buf, v)
				_, err := ln.Write(buf)
				if err != nil {
					fmt.Printf("client write error:%s\n", err.Error())
					return
				}
				fmt.Printf("client write:%d\n", v)
				time.Sleep(5 * time.Second)
			}
		}()
		wg.Wait()
		ch <- struct{}{}
	}(ln)
	<-ch
}
