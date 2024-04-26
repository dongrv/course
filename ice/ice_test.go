package ice

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
)

func Layout(rows, cols int) *Board {
	board := NewBoard(rows, cols)
	return board.Init()
}

func TestBoard_Init(t *testing.T) {
	Layout(6, 8).Print()
}

func TestBoard_CalcDirection(t *testing.T) {
	board := Layout(2, 2)
	board.Print()
	x, y := 0, 1
	block, err := board.Get(x, y)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if err = board.CalcDirection(block); err != nil {
		t.Error(err.Error())
		return
	}
	board.JSON(block)
}

func TestBoard_Click(t *testing.T) {
	board := Layout(6, 8)
	board.Print()

	board.Fishpond = &Fishpond{
		Fishes: []Fish{
			{Typ: BlueWhale},
			{Typ: BlueWhale},
			{Typ: Shark},
			{Typ: Crabs},
		},
	}
	board.Fishpond.Tidy()
}

func TestRecursive(t *testing.T) {
	output = true
	err := BatchRunRecursive()
	if err != nil {
		t.Fatalf("Result is fail:%s\n", err.Error())
	}
}

func BenchmarkRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := BatchRunRecursive(); err != nil {
			b.Fatalf("Result is fail:%s\n", err.Error())
		}
	}
}

var output bool

func BatchRunRecursive() error {
	funcs := []func() *Board{
		/*mockBoard1x1,
		mockBoard2x2,
		mockBoard2x3,*/
		mockBoard4x4,
		/*mockBoard5x5,
		mockBoard6x8,*/
	}
	for _, f := range funcs {
		board := f()
		if !board.Validate() {
			board.Print()
			return errors.New("validation failed")
		}
		if output {
			fmt.Print("初始布局\n")
			board.Print()
		}
		if !Recursive(board, board.Fishpond.NeedToPlace()) {
			board.Print()
			return errors.New("recursive failed")
		}
		if output {
			fmt.Print("计算结果\n")
			board.Print()
			board.JSON(nil)
		}

	}
	return nil
}

func mockBoard1x1() *Board {
	if output {
		fmt.Printf("1x1 board\n")
	}
	board := &Board{
		Rows: 1,
		Cols: 1,
		Fishpond: &Fishpond{
			Fishes: []Fish{
				{Typ: Blowfish},
			}},
	}
	board.Fishpond.Tidy()
	board.Scope = [][]Block{
		{NewBlock(0, 0)},
	}
	return board
}

func mockBoard2x2() *Board {
	if output {
		fmt.Printf("2x2 board\n")
	}
	board := &Board{
		Rows: 2,
		Cols: 2,
		Fishpond: &Fishpond{
			Fishes: []Fish{
				{Typ: Crabs},
				{Typ: Blowfish},
			}},
	}
	board.Fishpond.Tidy()
	board.Scope = [][]Block{
		{NewBlock(0, 0), NewBlock(1, 0).Cannot()},
		{NewBlock(0, 1), NewBlock(1, 1)},
	}
	return board
}

func mockBoard2x3() *Board {
	if output {
		fmt.Printf("2x3 board\n")
	}
	board := &Board{
		Rows: 2,
		Cols: 3,
		Fishpond: &Fishpond{
			Fishes: []Fish{
				{Typ: Crabs},
				{Typ: Crabs},
			}},
	}
	board.Fishpond.Tidy()
	board.Scope = [][]Block{
		{NewBlock(0, 0), NewBlock(1, 0), NewBlock(2, 0).Cannot()},
		{NewBlock(0, 1), NewBlock(1, 1), NewBlock(2, 1)},
	}
	return board
}

func mockBoard4x4() *Board {
	if output {
		fmt.Printf("4x4 board\n")
	}
	board := &Board{
		Rows: 4,
		Cols: 4,
		Fishpond: &Fishpond{
			Fishes: []Fish{
				{Typ: BlueWhale},
				{Typ: Crabs},
				{Typ: Shark},
				{Typ: Shark},
			}},
	}
	board.Fishpond.Tidy()
	board.Scope = [][]Block{
		{NewBlock(0, 0).Cannot(), NewBlock(1, 0).Cannot(), NewBlock(2, 0).Cannot(), NewBlock(3, 0)},
		{NewBlock(0, 1).Cannot(), NewBlock(1, 1), NewBlock(2, 1), NewBlock(3, 1)},
		{NewBlock(0, 2), NewBlock(1, 2), NewBlock(2, 2), NewBlock(3, 2)},
		{NewBlock(0, 3), NewBlock(1, 3), NewBlock(2, 3), NewBlock(3, 3)},
	}
	return board
}

func mockBoard5x5() *Board {
	if output {
		fmt.Printf("5x5 board\n")
	}
	board := &Board{
		Rows: 5,
		Cols: 5,
		Fishpond: &Fishpond{
			Fishes: []Fish{
				{Typ: Shark},
				{Typ: BlueWhale},
				{Typ: BlueWhale},
				{Typ: Crabs},
			}},
	}
	board.Fishpond.Tidy()
	board.Scope = [][]Block{
		{NewBlock(0, 0), NewBlock(1, 0), NewBlock(2, 0), NewBlock(3, 0), NewBlock(4, 0)},
		{NewBlock(0, 1), NewBlock(1, 1), NewBlock(2, 1), NewBlock(3, 1).Cannot(), NewBlock(4, 1)},
		{NewBlock(0, 2).Cannot(), NewBlock(1, 2).Cannot(), NewBlock(2, 2), NewBlock(3, 2).Cannot(), NewBlock(4, 2)},
		{NewBlock(0, 3).Cannot(), NewBlock(1, 3), NewBlock(2, 3).Cannot(), NewBlock(3, 3), NewBlock(4, 3)},
		{NewBlock(0, 4), NewBlock(1, 4), NewBlock(2, 4).Cannot(), NewBlock(3, 4), NewBlock(4, 4)},
	}
	return board
}

func mockBoard6x8() *Board {
	if output {
		fmt.Printf("6x8 board\n")
	}
	board := &Board{
		Rows: 6,
		Cols: 8,
		Fishpond: &Fishpond{
			Fishes: []Fish{
				{Typ: BlueWhale},
				{Typ: BlueWhale},
				{Typ: Shark},
				{Typ: Shark},
				{Typ: Blowfish},
				{Typ: Crabs},
				{Typ: Blowfish},
			}},
	}
	board.Fishpond.Tidy()
	board.Scope = [][]Block{
		{NewBlock(0, 0).Cannot(), NewBlock(1, 0).Cannot(), NewBlock(2, 0).Cannot(), NewBlock(3, 0), NewBlock(4, 0), NewBlock(5, 0), NewBlock(6, 0), NewBlock(7, 0)},
		{NewBlock(0, 1), NewBlock(1, 1), NewBlock(2, 1), NewBlock(3, 1), NewBlock(4, 1), NewBlock(5, 1), NewBlock(6, 1), NewBlock(7, 1)},
		{NewBlock(0, 2).Cannot(), NewBlock(1, 2), NewBlock(2, 2), NewBlock(3, 2), NewBlock(4, 2), NewBlock(5, 2), NewBlock(6, 2), NewBlock(7, 2)},
		{NewBlock(0, 3), NewBlock(1, 3).Cannot(), NewBlock(2, 3), NewBlock(3, 3), NewBlock(4, 3), NewBlock(5, 3), NewBlock(6, 3), NewBlock(7, 3)},
		{NewBlock(0, 4), NewBlock(1, 4), NewBlock(2, 4), NewBlock(3, 4), NewBlock(4, 4), NewBlock(5, 4), NewBlock(6, 4), NewBlock(7, 4)},
		{NewBlock(0, 5), NewBlock(1, 5).Cannot(), NewBlock(2, 5), NewBlock(3, 5), NewBlock(4, 5), NewBlock(5, 5).Cannot(), NewBlock(6, 5), NewBlock(7, 5)},
	}
	return board
}

func TestRecursive2(t *testing.T) {
	var (
		success int
		fail    int
	)
	for i := 0; i < 100000; i++ {
		fishes := []Fish{
			{Typ: BlueWhale},
			{Typ: Crabs},
			{Typ: BlueWhale},
			{Typ: Shark},
			{Typ: Shark},
			{Typ: Blowfish},
			{Typ: BlueWhale},
		} // 超过10条鱼，递归深度导致的开销非常大
		err := MockRandomClick(6, 8, fishes, 6)
		if err != nil {
			fail++
			continue
		}
		success++
	}
	t.Logf("success:%d, fail:%d", success, fail)
}

func MockRandomClick(rows, cols int, fishes []Fish, maxRandomNum int) error {
	board := Layout(rows, cols).Init()
	if !board.Validate() {
		return errors.New("validate failed")
	}
	board.Fishpond = NewFishpond(fishes)
	store := map[[2]int]struct{}{}
	counter := 0
	for {
		x := rand.Intn(cols)
		y := rand.Intn(rows)
		if _, ok := store[[2]int{x, y}]; ok {
			continue
		}
		store[[2]int{x, y}] = struct{}{}
		counter++
		if counter == maxRandomNum {
			break
		}
	}
	// 设置空占位
	for coord := range store {
		block, err := board.Get(coord[0], coord[1])
		if err != nil {
			return err
		}
		*block = (*block).Cannot()
	}
	if !Recursive(board, board.Fishpond.NeedToPlace()) {
		board.ScopeJSON()
		board.JSON(board.Fishpond)
		return errors.New("recursive failed")
	}
	return nil
}

func TestReverse(t *testing.T) {
	list := []Block{NewBlock(0, 5), NewBlock(1, 5)}
	list = Reverse(list)
	t.Logf("list: %v", list)
}
