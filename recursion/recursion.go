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
// 递归的缺点：
// - 每次递推调用，都会增加一层栈帧，每次返回都会释放一层栈帧，如果程序递归层次很深，会占用很多堆栈空间，甚至会导致栈溢出 stack overflow
// - 递归在一些场景下是很高效且简单的，比如有固定规律的数学运算、路径方案，但递归程序效率高和低都是基于实际场景，要开发人员能够甄别是否适合使用

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

// 回溯
// 回溯算法实际上一个类似枚举的搜索尝试过程，主要是在搜索尝试过程中寻找问题的解，当发现已不满足求解条件时，就“回溯”返回，尝试别的路径。
// 回溯法是一种选优搜索法，按选优条件向前搜索，以达到目标。但当探索到某一步时，发现原先选择并不优或达不到目标，就退回一步重新选择，这种走不通就退回再走的技术为回溯法，
// 而满足回溯条件的某个状态的点称为“回溯点”

/*

				  root0
                 /     \
		       /        \
             /           \
           n1            n2
          /  \         /    \
        n3   n4       n5    n6
       / \  / \     / \    /  \
     n7 n8 n9 n10 n11 n12 n13 n14


void backtracking(参数) {
    if (终止条件) {
        存放结果;
        return;
    }
    for (选择：本层集合中元素（树中节点孩子的数量就是集合的大小）) {
        处理节点;
        backtracking(路径，选择列表); // 递归
        回溯，撤销处理结果
    }
}

*/
