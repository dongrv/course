package tdd

import (
	"math"
	"strconv"
)

// 测试驱动开发（Test-Driven Development） - TDD
// 拓展阅读：https://zhuanlan.zhihu.com/p/404832754
// 它要求在编写某个功能的代码之前先编写测试代码，然后只编写使测试通过的功能代码，通过测试来推动整个开发的进行。这有助于编写简洁可用和高质量的代码，并加速开发过程。

// AddOne 累加1
func AddOne(x int) int {
	return x + 1
}

// Transform 字符串转数字
func Transform(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		return -1
	}
	return i
}

// Distance 计算两个坐标之间的距离
func Distance(x1, y1, x2, y2 float64) float64 {
	v := math.Abs(x1 - x2)
	o := math.Abs(y1 - y2)
	return math.Sqrt(v*v + o*o)
}
