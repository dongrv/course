package learn_aceld

// linux time 工具
// time go run main.go
// real	0m0.015s 从程序开始到结束，实际度过的时间；
// user	0m0.001s 程序在用户态度过的时间；
// sys	0m0.006s 程序在内核态度过的时间。
// 一般情况下 real >= user + sys，因为系统还有其它进程(切换其他进程中间对于本进程会有空白期)。

// /usr/bin/time -v  go run

// Command being timed: "go run ."
// 更详细的结果：
// User time (seconds): 0.01
// System time (seconds): 0.17
// Percent of CPU this job got: 56%
// Elapsed (wall clock) time (h:mm:ss or m:ss): 0:00.34
// Average shared text size (kbytes): 0
// Average unshared data size (kbytes): 0
// Average stack size (kbytes): 0
// Average total size (kbytes): 0
// Maximum resident set size (kbytes): 23424
// Average resident set size (kbytes): 0
// Major (requiring I/O) page faults: 226
// Minor (reclaiming a frame) page faults: 11580
// Voluntary context switches: 2544
// Involuntary context switches: 70
// Swaps: 0
// File system inputs: 53832
// File system outputs: 1712
// Socket messages sent: 0
// Socket messages received: 0
// Signals delivered: 0
// Page size (bytes): 4096
// Exit status: 0
//
// 更丰富的信息：
// CPU占用率；
// 内存使用情况；
// Page Fault 情况；
// 进程切换情况；
// 文件系统IO；
// Socket 使用情况；
// ……

/*
package main

import (
    "log"
    "runtime"
    "time"
)

func test() {
    //slice 会动态扩容，用slice来做堆内存申请
    container := make([]int, 8)

    log.Println(" ===> loop begin.")
    for i := 0; i < 32*1000*1000; i++ {
        container = append(container, i)
    }
    log.Println(" ===> loop end.")
}

func main() {
    log.Println("Start.")

    test()

    log.Println("force gc.")
    runtime.GC() //强制调用gc回收

    log.Println("Done.")

    time.Sleep(3600 * time.Second) //睡眠，保持程序不退出
}
*/

// go build -o snippet_mem && ./snippet_mem

// 查看一个运行中的进程的资源占用情况
// top -p $(pidof snippet_mem)
// 显示大量内存没有被释放

// 跟踪打印垃圾回收器信息
// GODEBUG='gctrace=1' ./snippet_mem
