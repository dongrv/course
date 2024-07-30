package train

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

// https://xie.infoq.cn/article/495308325c34c996632275b7c

func Div(a, b int32) float64 {
	return float64(a) / float64(b)
}

func SafeGet(list []int32, n int) int32 {
	if len(list) <= n {
		return 0
	}
	return list[n]
}

// FindLess 找出小于目标值的数据
func FindLess(source []int8, target int8) []int8 {
	var result []int8
	for _, i := range source {
		if i < target {
			result = append(result, i)
		}
	}
	return result
}

// Random 从切片随机一个值
func Random(list []int8) int8 {
	return list[rand.Intn(len(list))]
}

// Match 基于需求匹配一个值
func Match(value []int8, bound int8) int8 {
	return Random(FindLess(value, bound))
}

func CausePanic() {
	// 空指针
	type Foo struct{ Id int }
	dict := map[int32]*Foo{1: {Id: 100}}
	if _, ok := dict[0]; ok {
		fmt.Printf("VIP Id:%d\n", dict[0].Id)
	}
	// Map读写并发导致系统崩溃：concurrent map iteration and map write
	// 此处假设dict元素很多
	go func() {
		for k, v := range dict {
			fmt.Printf("k:%d -> v:%+v\n", k, v)
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			dict[int32(i)] = &Foo{Id: i}
		}
	}()

	// 切片越界
	slice := []int8{1, 2, 3, 4, 5, 6}
	n := 6
	if len(slice) > n {
		fmt.Printf("index:%d -> value:%d\n", n, slice[n])
	}
	// 协程死锁
	var mu sync.RWMutex
	go func() {
		mu.Lock()
		// 缺少锁释放
		time.Sleep(10 * time.Second)
		// todo something
		return
	}()
	go func() {
		mu.RLock()
		defer mu.RUnlock()
		// todo something
	}()

	// I/O
	// · 操作都是不安全的，需要检测返回错误
	// · 必须关闭资源
	// · 耗时长的操作提前评估和重新决策方案
	file, err := os.Open("./file.txt")
	if err != nil {
		//return err
		return
	}
	defer file.Close()

	// 指针参数
	type Number struct{ num int }
	var num *Number

	go func(value *Number) {
		println(value.num) // panic:value is nil
	}(num)

	// 没有明确退出条件的循环
	for {
		// todo something
	}

}

type Bar struct {
	mu     sync.RWMutex
	Number int
}

func (b *Bar) Get() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Number
}
func (b *Bar) Counter() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Number++
	return b.Get()
}

type Coord struct {
	Country  string
	Province string
	Region   *Region
}

type Region struct {
	Zip int64
}

func Analyze() {
	coord := new(Coord)
	fmt.Printf("coord addr:%p\n", &coord)
	fmt.Printf("coord pointer value:%p\n", coord)
	fmt.Printf("coord.Region addr:%p\n", &coord.Region)
	fmt.Printf("coord.Region pointer value:%p\n", coord.Region)
	fmt.Printf("coord.Region object value:%v\n", *coord.Region)
}

func AnalyzeSlice() {
	// 情景一
	buf := make([]byte, 1024)
	buf[0] = 'a'
	fmt.Printf("buf addr:%p\n", &buf)
	fmt.Printf("buf reference addr:%p\n", buf)
	fmt.Printf("buf[0] addr:%p\n", &buf[0])
	buf = growSlice(buf)
	fmt.Printf("after grow buf addr:%p\n", &buf)
	fmt.Printf("after grow buf reference addr:%p\n", buf)
	fmt.Printf("after grow buf[0] addr:%p\n", &buf[0])
	println(strings.Repeat("=", 100))
	// 情景二
	buf2 := make([]byte, 512, 1024)
	fmt.Printf("buf2 addr:%p\n", &buf2)
	fmt.Printf("buf2 reference addr:%p\n", buf2)
	fmt.Printf("buf2[0] addr:%p\n", &buf2[0])
	buf2 = growSlice(buf2)
	fmt.Printf("after grow buf addr:%p\n", &buf2)
	fmt.Printf("after grow buf reference addr:%p\n", buf2)
	fmt.Printf("after grow buf[0] addr:%p\n", &buf2[0])
}

func growSlice(buf []byte) []byte {
	buf = append(buf, 'b')
	return buf
}
