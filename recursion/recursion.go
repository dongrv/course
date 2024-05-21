package recursion

// 斐波那契数列

// f(0) = 0
// f(1) = 1
// f(2) = f(2-1) + f(2 -2) = f(1) + f(0)
//... 从 n>=2 开始，每一项的结果都是前两项之和

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// 阶乘

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
