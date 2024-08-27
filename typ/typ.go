package typ

import (
	"math"
)

func Assign() {
	a, b, c, d := 0b111, 0o123, 0x64, 1_000_000_000
	println(a, b, c, d)

	// x := nil // 报错：默认零值，没有类型，不能作为初始值供编译器断言，不同类型的nil值不能比较
}

func MathLog10() {
	a := 123456
	b := 123
	c := 103450
	println(
		math.Log10(float64(a)),
		math.Log10(float64(b)),
		math.Log10(float64(c)),
	)

}
