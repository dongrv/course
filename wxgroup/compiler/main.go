package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const numElements = 1000000

var numWorkers = runtime.NumCPU()

// 模拟一个计算密集型函数
func compute(n int) int {
	var result int
	for i := 0; i < n; i++ {
		result += i * i
	}
	return result
}

// 使用并发优化计算
func parallelCompute(data []int) (total int) {
	var wg sync.WaitGroup
	totals := make([]int, numWorkers)

	chunkSize := numElements / numWorkers
	for i := 0; i < numWorkers; i++ {
		from := i * chunkSize
		to := from + chunkSize
		if i == numWorkers-1 {
			to = numElements
		}
		wg.Add(1)
		go func(from, to int) {
			defer wg.Done()
			for j := from; j < to; j++ {
				totals[i] += compute(data[j])
			}
		}(from, to)
	}
	wg.Wait()

	for _, t := range totals {
		total += t
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	data := make([]int, numElements)
	for i := range data {
		data[i] = rand.Intn(100000)
	}

	start := time.Now()
	sequentialResult := 0
	for i := 0; i < numElements; i++ {
		sequentialResult += compute(data[i])
	}
	fmt.Printf("Sequential Time: %s Total: %d\n", time.Since(start), sequentialResult)

	start = time.Now()
	parallelResult := parallelCompute(data)
	fmt.Printf("Parallel Time: %s Total: %d\n", time.Since(start), parallelResult)
}
