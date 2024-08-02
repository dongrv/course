package main

import (
	"fmt"
	"math/rand"
)

func generate8191() {
	nums := make([]int, 8191) // < 64KB
	for i := 0; i < 8191; i++ {
		nums[i] = rand.Int()
	}
}

func generate8193() {
	nums := make([]int, 8193) // > 64KB
	for i := 0; i < 8193; i++ {
		nums[i] = rand.Int()
	}
}

func generate(n int) {
	nums := make([]int, n) // 不确定大小
	for i := 0; i < n; i++ {
		nums[i] = rand.Int()
	}
}

type Demo struct{}

func NewDemo() *Demo {
	return &Demo{}
}

func main() {
	// 指针
	demo := NewDemo()
	_ = demo
	// interface{} 参数
	fmt.Println(Demo{})

	// 内存大小
	generate8191()
	generate8193()
	generate(1)

	// 闭包
	in := Increase()
	fmt.Println(in()) // 1
	fmt.Println(in()) // 2
}

func Increase() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}
