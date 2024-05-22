package recursion

import (
	"fmt"
	"strings"
)

// 递归是一种解决问题的有效方法，在递归过程中，函数将自身作为子例程调用。
// 递归的基本思想是某个函数直接或者间接地调用自身，这样原问题的求解就转换为了许多性质相同但是规模更小的子问题。
// 可以认为是作为程序单元，“我”调用“我”自己，不断将顶层问题拆解，直到遇到条件的临界点返回。
// 三大要素：
// - 明确函数的功能
// - 寻找递归结束的条件
// - 找出函数的等价关系

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

// 0！ = 1
// 1! = 1 * 0!
// 2! = 2 * 1 * 0!
// ...
// 10! = 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1 * 0!

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// 棋盘寻路

// Coord 坐标
type Coord struct {
	X, Y   int
	Placed bool // 是否已放置
}

func (c Coord) Clone() Coord {
	return c
}

// ChessBoard 棋盘
type ChessBoard struct {
	Scope    [3][3]Coord
	From, To Coord
}

func NewChessBoard() *ChessBoard {
	return &ChessBoard{}
}

func (board *ChessBoard) Init() *ChessBoard {
	for y, coords := range board.Scope {
		for x, coord := range coords {
			coord.X, coord.Y = x, y
		}
	}
	return board
}

func (board *ChessBoard) FromTo(from, to Coord) *ChessBoard {
	board.From, board.To = from, to
	return board
}

const (
	Up = iota
	Down
	Left
	Right
)

var directions = [4]int{Up, Down, Left, Right}

func Search(board *ChessBoard, coord Coord) bool {
	if coord.Placed {
		return false
	}
	if (board.From == board.To) || (board.To == coord) {
		return true
	}
	if 0 > coord.Y || coord.Y > len(board.Scope[0])-1 {
		return false // 突破上下边界
	}
	if 0 > coord.X || coord.X > len(board.Scope[coord.Y])-1 {
		return false // 突破左右边界
	}
	// 遍历四个方向
	for _, direction := range directions {
		switch direction {
		case Up:
			board.Scope[coord.Y][coord.X].Placed = true
			newCoord := coord
			newCoord.Y--
			if Search(board, newCoord) {
				return true
			}
			board.Scope[coord.Y][coord.X].Placed = false
			return false
		case Down:
			board.Scope[coord.Y][coord.X].Placed = true
			newCoord := coord
			newCoord.Y++
			if Search(board, newCoord) {
				return true
			}
			board.Scope[coord.Y][coord.X].Placed = false
			return false
		case Left:
			board.Scope[coord.Y][coord.X].Placed = true
			newCoord := coord
			newCoord.X--
			if Search(board, newCoord) {
				return true
			}
			board.Scope[coord.Y][coord.X].Placed = false
			return false
		case Right:
			board.Scope[coord.Y][coord.X].Placed = true
			newCoord := coord
			newCoord.X++
			if Search(board, newCoord) {
				return true
			}
			board.Scope[coord.Y][coord.X].Placed = false
			return false
		}
	}
	return false
}

// Print 打印棋盘
func (board *ChessBoard) Print() *ChessBoard {
	println(strings.Repeat("=", 20))
	for _, rows := range board.Scope {
		for _, coord := range rows {
			fmt.Printf("(%d, %d)\t", coord.X, coord.Y)
		}
		println()
	}
	return board
}
