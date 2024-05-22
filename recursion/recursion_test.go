package recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFib(t *testing.T) {
	assert.Equal(t, 0, Fib(0))
	assert.Equal(t, 1, Fib(1))
	assert.Equal(t, 1, Fib(2))
	assert.Equal(t, 2, Fib(3))
	assert.Equal(t, 5, Fib(5))
	assert.Equal(t, 55, Fib(10))
}

/*

TODO Fib(5)的展开式：

先递推 -> 再回归

Fib(5)
= Fib(4) + Fib(3)
= (Fib(3) + Fib(2)) + (Fib(2) + Fib(1))
= (Fib(2) + Fib(1) + (Fib(1) + Fib(0))) + ((Fib(1) + Fib(0)) + Fib(1))
= (((Fib(1) + Fib(0)) + Fib(1)) + (Fib(1) + Fib(0))) + ((Fib(1) + Fib(0)) + Fib(1))
= (((1 + 0) + 1) + (1 + 0)) + ((1 + 0) + 1)
= (((1) + 1) + (1)) + ((1) + 1)
= ((2) + (1)) + ((1) + 1)
= (3)) + (2)
= 5

*/

func TestFactorial(t *testing.T) {
	assert.Equal(t, 1, Factorial(0))
	assert.Equal(t, 6, Factorial(3))
	assert.Equal(t, 3628800, Factorial(10))
}

func TestSearch(t *testing.T) {
	start, end := Coord{X: 0, Y: 0}, Coord{X: 2, Y: 2}
	board := NewChessBoard().
		Init().
		Print().
		FromTo(start, end)

	assert.True(t, Search(board, start))

}
