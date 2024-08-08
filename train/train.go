package train

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"
	"unsafe"
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
	fmt.Printf("in grow buf addr:%p\n", &buf)
	fmt.Printf("in grow buf reference addr:%p\n", buf)
	return buf
}

func SliceReference() {
	s := []int{123, 456, 789}
	println()
	fmt.Printf("&s=%p\ns=%p\n\n", &s, s)
	fmt.Printf("&s[0]=%p\n&s[1]=%p\n&s[1]=%p\n\n", &s[0], &s[1], &s[2])
}

func AccessUnderlyingSlice() {
	s := make([]int, 0, 5)
	s = append(s, 123, 456, 789)
	headerPtr := (*[3]uintptr)(unsafe.Pointer(&s))
	arrayPtr := *(*int)(unsafe.Pointer(&headerPtr[0])) // 底层数组指针
	length := *(*int)(unsafe.Pointer(&headerPtr[1]))   // 长度
	capacity := *(*int)(unsafe.Pointer(&headerPtr[2])) // 容量
	fmt.Println("当前s的底层值：")
	fmt.Printf("type slice struct{\n\tarray \t= 0x%x\n\tlen \t= %d\n\tcap \t= %d\n\n}\n",
		arrayPtr, length, capacity)
	fmt.Println("访问s.array的值：")
	for i := 0; i < length; i++ {
		fmt.Printf("&s[%d]=%d\n", i, *(*int)(unsafe.Pointer(uintptr(arrayPtr) + uintptr(i)*unsafe.Sizeof(arrayPtr))))
	}
}

func ParseFuncStack() {
	var (
		a int   // 0
		b []int // slice{array:nil, len:0, cap:0}
	)
	b = append(b, 1, 2, 3, 4, 5) // slice{array:&[6]{1, 2, 3, 4, 5, NULL}, len:5, cap:6}
	fmt.Printf("a=%d\nb.array=%v len(b)=%d cap(b)=%d\n", a, b, len(b), cap(b))
}

// 逃逸分析

type Role struct {
	Id   int32
	Play string
}

func PointerAndEscape() map[string]*Role {
	roles := make(map[string]*Role)
	roles["Tom"] = &Role{Id: 0, Play: "Tiger"}
	roles["Lucy"] = &Role{Id: 0, Play: "Lion"}
	roles["Mark"] = &Role{Id: 0, Play: "Leopard"}
	//roles := map[string]*Role{
	//	"Tom":  {Id: 0, Play: "Tiger"},
	//	"Lucy": {Id: 1, Play: "Lion"},
	//	"Mark": {Id: 2, Play: "Leopard"},
	//}
	changeRole(&(*roles["Tom"]), "Bear")
	fmt.Printf("tom value pointer:%p\n", roles["Tom"])
	role := *roles["Lucy"]
	changeRole(&role, "Wolf")
	fmt.Printf("tom value pointer:%p\n", roles["Lucy"])
	return roles
}

func changeRole(r *Role, play string) {
	println(strings.Repeat("=", 100))
	r.Play = play
	fmt.Printf("in func %d value pointer:%p\n", r.Id, r)
}

type Gopher struct{}

func (g *Gopher) Code()          {}
func (g *Gopher) Name() string   { return "go" }
func (g *Gopher) Author() string { return "google" }

// 反射方法名按照字典排序a-z

func ReflectMethod() {
	g := &Gopher{}
	typ := reflect.TypeOf(g)
	for i := 0; i < typ.NumMethod(); i++ {
		fmt.Printf("methed:%s\n", typ.Method(i).Name)
	}
}

// DeleteSlice 删除指定索引前的元素
func DeleteSlice() {
	var a = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}
	i := 1
	if i < len(a)-1 {
		copy(a[i:], a[i+1:])
	}
	a[len(a)-1] = nil // 多出一个尾部元素，空余位置置空有助于垃圾回收
	a = a[:len(a)-1]
	fmt.Printf("%v", a)
}
