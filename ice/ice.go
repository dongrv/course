package ice

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

var (
	ErrIceBlockBroken   = errors.New("ice block broken")    // 冰块已被击碎过
	ErrIndexOutOfBounds = errors.New("index out of bounds") // 访问越界
	ErrNoMoreFish       = errors.New("no more fish")        // 没有更多鱼
	ErrBlockTaken       = errors.New("block taken")         // 位置被占
	ErrInvalidDirection = errors.New("invalid direction")   // 无效的方向参数
	ErrInvalidParameter = errors.New("invalid parameter")   // 无效参数
)

type Direction int8

const (
	Left Direction = iota
	Right
	Orthogonal // 横向
	Up
	Down
	Vertical // 竖向
)

var Directions = []Direction{Left, Right, Orthogonal, Up, Down, Vertical}

var directionMap = map[Direction]string{
	Up:         "Up",
	Down:       "Down",
	Vertical:   "Vertical",
	Left:       "Left",
	Right:      "Right",
	Orthogonal: "Orthogonal",
}

// 横向/水平
func (direction Direction) orthogonal() bool {
	return direction == Left || direction == Right || direction == Orthogonal
}

// 竖向/垂直
func (direction Direction) vertical() bool {
	return direction == Up || direction == Down || direction == Vertical
}

func (direction Direction) String() string {
	return directionMap[direction]
}

// Meta 方位基本信息
type Meta struct {
	Towards Direction // 朝向
	Length  int
	From    Coordinate
	To      Coordinate
	Fish    FishType
}

// Coordinate 坐标
type Coordinate struct {
	X int
	Y int
}

// Block 冰块
type Block struct {
	CanUse      bool       // 当前块是否可用
	Coord       Coordinate // 当前坐标
	Metas       [6]Meta    // 方位和举例
	Direction   Direction  // 方向索引，默认 -1
	Fish        FishType   // 鱼的类型
	FishToIndex int        // 对应的索引
	Clicked     bool       // 是否点击过当前冰块
}

func NewBlock(x, y int) *Block {
	return &Block{CanUse: true, Coord: Coordinate{X: x, Y: y}, Direction: -1}
}

func (b *Block) Can() *Block {
	b.CanUse = true
	b.Fish = None
	return b
}

func (b *Block) Cannot() *Block {
	b.CanUse = false
	b.Fish = Blank
	return b
}

// TakeOver 当前快被占领
func (b *Block) TakeOver(fish Fish, direction Direction) {
	b.CanUse, b.Fish, b.Direction, b.FishToIndex = false, fish.Typ, direction, fish.Index
}

// Cancel 取消之前的设置
func (b *Block) Cancel() {
	b.CanUse, b.Fish, b.Direction, b.FishToIndex = true, 0, -1, 0
}

//+----------------------------------------------------------------+
//+					         Board Handle	                       +
//+----------------------------------------------------------------+

type Board struct {
	Rows     int       // 行数:y轴
	Cols     int       // 列数:x轴
	Scope    [][]Block // 分布：二维数组
	Fishpond *Fishpond // 有多少鱼需要放置
}

func NewBoard(rows, cols int) *Board {
	return &Board{Rows: rows, Cols: cols}
}

func (board *Board) Init() *Board {
	if board.Rows == 0 || board.Cols == 0 {
		return &Board{} // 参数错误
	}
	board.Scope = make([][]Block, 0, board.Rows)
	for i := 0; i < board.Rows; i++ {
		row := make([]Block, 0, board.Cols)
		for j := 0; j < board.Cols; j++ {
			row = append(row, Block{
				CanUse:    true,
				Coord:     Coordinate{j, i},
				Direction: -1,
			})
		}
		board.Scope = append(board.Scope, row)
	}
	return board
}

// Validate 校验参数
func (board *Board) Validate() bool {
	if len(board.Scope) == 0 {
		return false
	}
	return len(board.Scope) == board.Rows && len(board.Scope[0]) == board.Cols
}

// CanPut 当前坐标是否可放置
func (board *Board) CanPut(x, y int) bool {
	if board.Rows <= y || board.Cols <= x {
		return false
	}
	return board.Scope[y][x].CanUse
}

func (board *Board) Cancel(block Block) error {
	coord := block.Coord
	meta := block.Metas[block.Direction]
	board.Scope[coord.Y][coord.X] = *NewBlock(block.Coord.X, block.Coord.Y) // 先放置当前坐标
	if meta.From == meta.To {
		return nil // 只需要放置当前坐标
	}
	length := block.Fish.Value()
	if meta.Towards.orthogonal() {
		// 水平摆放
		start := min(meta.From.X, meta.To.X)
		end := max(meta.From.X, meta.To.X)
		if (end - start + 1) < length { // +1 要包含所有模块
			return ErrInvalidParameter
		}
		for offset := start; offset < start+length; offset++ {
			block.Coord.X, block.Coord.Y = offset, meta.From.Y
			board.Scope[meta.From.Y][offset] = *NewBlock(block.Coord.X, block.Coord.Y)
		}
	} else {
		// 垂直摆放
		start := min(meta.From.Y, meta.To.Y)
		end := max(meta.From.Y, meta.To.Y)
		if (end - start + 1) < length {
			return ErrInvalidParameter
		}
		for offset := start; offset < start+length; offset++ {
			block.Coord.X, block.Coord.Y = meta.From.X, offset
			board.Scope[offset][meta.From.X] = *NewBlock(block.Coord.X, block.Coord.Y)
		}
	}
	return nil
}

// TakeOver 占领棋盘
func (board *Board) TakeOver(block Block) ([]Block, error) {
	if !board.CanPut(block.Coord.X, block.Coord.Y) {
		return nil, ErrBlockTaken
	}
	if block.Direction < 0 {
		return nil, ErrInvalidDirection
	}
	block.Clicked = false // 默认未点击过
	taken := make([]Block, 0, 3)
	coord := block.Coord
	meta := block.Metas[block.Direction]
	board.Scope[coord.Y][coord.X] = block // 先放置当前坐标
	if meta.From == meta.To {
		return []Block{block}, nil // 只需要放置当前坐标
	}
	length := block.Fish.Value()
	if meta.Towards.orthogonal() {
		// 水平摆放
		start := min(meta.From.X, meta.To.X)
		end := max(meta.From.X, meta.To.X)
		if (end - start + 1) < length {
			return nil, ErrInvalidParameter
		}
		var cursor int // 游标，用于定位开始索引

		if rand.Intn(2)%2 == 0 { // 增加随机性
			if block.Coord.X+length <= end {
				cursor = block.Coord.X
			} else if block.Coord.X-length+1 >= start {
				cursor = block.Coord.X - length + 1
			} else {
				cursor = start
			}
		} else {
			if block.Coord.X-length+1 >= start {
				cursor = block.Coord.X - length + 1
			} else if block.Coord.X+length <= end {
				cursor = block.Coord.X
			} else {
				cursor = start
			}
		}

		for offset := cursor; offset < cursor+length; offset++ {
			block.Coord.X, block.Coord.Y = offset, meta.From.Y
			board.Scope[meta.From.Y][offset] = block
			taken = append(taken, block)
		}
	} else {
		// 垂直摆放
		start := min(meta.From.Y, meta.To.Y)
		end := max(meta.From.Y, meta.To.Y)
		if (end - start + 1) < length {
			return nil, ErrInvalidParameter
		}
		var cursor int           // 游标，用于定位开始索引
		if rand.Intn(2)%2 == 0 { // 增加随机性
			if block.Coord.Y+length <= end {
				cursor = block.Coord.Y
			} else if block.Coord.Y-length+1 >= start {
				cursor = block.Coord.Y - length + 1
			} else {
				cursor = start
			}
		} else {
			if block.Coord.Y-length+1 >= start {
				cursor = block.Coord.Y - length + 1
			} else if block.Coord.Y+length <= end {
				cursor = block.Coord.Y
			} else {
				cursor = start
			}
		}
		for offset := cursor; offset < cursor+length; offset++ {
			block.Coord.X, block.Coord.Y = meta.From.X, offset
			board.Scope[offset][meta.From.X] = block
			taken = append(taken, block)
		}
	}
	return taken, nil
}

// Get 根据坐标获取元素对应指针
func (board *Board) Get(x, y int) (*Block, error) {
	if x >= board.Cols || y >= board.Rows {
		return nil, ErrIndexOutOfBounds
	}
	return &board.Scope[y][x], nil
}

// CalcDirection 计算各个方向
func (board *Board) CalcDirection(b *Block) error {
	if len(board.Scope) <= b.Coord.Y {
		return ErrIndexOutOfBounds
	}
	if len(board.Scope[b.Coord.Y]) <= b.Coord.X {
		return ErrIndexOutOfBounds
	}
	if !board.Scope[b.Coord.Y][b.Coord.X].Fish.None() {
		return ErrBlockTaken // 当前位置已经被占了
	}
	for _, direction := range Directions {
		var length int
		switch direction {
		case Up:
			if b.Coord.Y == 0 {
				b.Metas[Up] = Meta{Towards: Up, Length: 1, From: b.Coord, To: b.Coord} // 第一行
				continue
			}
			var Y int
			for y := b.Coord.Y; y >= 0; y-- {
				if !board.stepOK(b.Coord.X, y, &length) {
					break
				}
				Y = y
			}
			b.Metas[Up] = Meta{
				Towards: Up,
				Length:  length,
				From:    b.Coord,
				To:      Coordinate{X: b.Coord.X, Y: Y},
			}
		case Down:
			if b.Coord.Y == board.Rows-1 {
				b.Metas[Down] = Meta{Towards: Down, Length: 1, From: b.Coord, To: b.Coord} // 最后一行
				continue
			}
			var Y int
			for y := b.Coord.Y; y < board.Rows; y++ {
				if !board.stepOK(b.Coord.X, y, &length) {
					break
				}
				Y = y
			}
			b.Metas[Down] = Meta{
				Towards: Down,
				Length:  length,
				From:    b.Coord,
				To:      Coordinate{X: b.Coord.X, Y: Y},
			}
		case Left:
			if b.Coord.X == 0 {
				b.Metas[Left] = Meta{Towards: Left, Length: 1, From: b.Coord, To: b.Coord} // 第一行
				continue
			}
			var X int
			for x := b.Coord.X; x >= 0; x-- {
				if !board.stepOK(x, b.Coord.Y, &length) {
					break
				}
				X = x
			}
			b.Metas[Left] = Meta{
				Towards: Left,
				Length:  length,
				From:    b.Coord,
				To:      Coordinate{X: X, Y: b.Coord.Y},
			}
		case Right:
			if b.Coord.X == board.Cols-1 {
				b.Metas[Right] = Meta{Towards: Right, Length: 1, From: b.Coord, To: b.Coord} // 最后一行
				continue
			}
			var X int
			for x := b.Coord.X; x < board.Cols; x++ {
				if !board.stepOK(x, b.Coord.Y, &length) {
					break
				}
				X = x
			}
			b.Metas[Right] = Meta{
				Towards: Right,
				Length:  length,
				From:    b.Coord,
				To:      Coordinate{X: X, Y: b.Coord.Y},
			}
		case Orthogonal:
			b.Metas[Orthogonal] = Meta{
				Towards: Orthogonal,
				Length:  b.Metas[Left].Length + b.Metas[Right].Length - 1,
				From:    b.Metas[Left].To,
				To:      b.Metas[Right].To,
			}
		case Vertical:
			b.Metas[Vertical] = Meta{
				Towards: Vertical,
				Length:  b.Metas[Up].Length + b.Metas[Down].Length - 1,
				From:    b.Metas[Up].To,
				To:      b.Metas[Down].To,
			}
		}
	}
	return nil
}

func (board *Board) stepOK(x, y int, length *int) bool {
	if !board.Scope[y][x].Fish.None() {
		return false
	} else {
		*length++
		return true
	}
}

// GetUsableBlocks 获取当前可用的块
func (board *Board) GetUsableBlocks() ([]Block, error) {
	if board.Rows == 0 || board.Cols == 0 || len(board.Scope) == 0 {
		return nil, ErrInvalidParameter
	}
	result := make([]Block, 0, board.Rows*board.Cols/2)
	for _, blocks := range board.Scope {
		for _, block := range blocks {
			if block.Fish.None() {
				if err := board.CalcDirection(&block); err != nil { // 计算各方向容量
					return nil, err
				}
				result = append(result, block)
			}
		}
	}
	return result, nil
}

// Print 打印棋盘
func (board *Board) Print() {
	if len(board.Scope) == 0 {
		return
	}
	fmt.Printf("%s\t\n", strings.Repeat("=", (board.Rows)*5-1))
	for _, blocks := range board.Scope {
		for _, block := range blocks {
			fmt.Printf("%d\t", block.Fish.Value())
		}
		println()
	}
	fmt.Printf("%s\t\n", strings.Repeat("=", (board.Rows)*5-1))
}

// ScopeJSON 打印json字符串
func (board *Board) ScopeJSON() string {
	bs, _ := json.Marshal(board.Scope)
	return string(bs)
}

// JSON 打印json字符串
func (board *Board) JSON(target interface{}) {
	if target != nil {
		bs, _ := json.Marshal(target)
		fmt.Printf("%s\n", string(bs))
		return
	}
	if len(board.Scope) == 0 {
		return
	}
	bs, _ := json.Marshal(board)
	fmt.Printf("%s\n", string(bs))
}

func Recursive(board *Board, fishes []Fish) bool {
	if len(fishes) == 0 {
		return true
	}
	blocks, err := board.GetUsableBlocks()
	if err != nil {
		fmt.Printf("Recursive err:%s", err)
		return false
	}
	if len(blocks) == 0 {
		return false
	}
	fish := fishes[0]
	fishes = fishes[1:]
	if !CanPlace(blocks, fish) {
		return false
	}
	for _, block := range blocks {
		// 尝试横向放置
		for _, direction := range Directions[:3] {
			if fish.Typ.Value() <= block.Metas[direction].Length {
				// 占位
				block.TakeOver(fish, direction)
				if _, err := board.TakeOver(block); err != nil {
					board.Print()
					fmt.Printf("take over err:%s\n", err.Error())
					return false
				}
				// 进入下一层
				if Recursive(board, fishes) {
					board.Fishpond.UpdateState(fish.Index, true)
					return true
				}
				// 撤销
				if err := board.Cancel(block); err != nil {
					board.Print()
					fmt.Printf("cancel board err:%s\n", err.Error())
				}
				block.Cancel()
			}
		}
		// 尝试竖向放置
		for _, direction := range Directions[3:] {
			if fish.Typ.Value() <= block.Metas[direction].Length {
				// 占位
				block.TakeOver(fish, direction)
				if _, err := board.TakeOver(block); err != nil {
					board.Print()
					fmt.Printf("take over err:%s\n", err.Error())
					return false
				}
				// 进入下一层
				if Recursive(board, fishes) {
					board.Fishpond.UpdateState(fish.Index, false)
					return true
				}
				// 撤销
				if err := board.Cancel(block); err != nil {
					board.Print()
					fmt.Printf("cancel board err:%s\n", err.Error())
				}
				block.Cancel()
			}
		}
	}
	return false
}

// CanPlace 至少存在一种情况可以放置当前鱼
func CanPlace(blocks []Block, fish Fish) bool {
	for _, block := range blocks {
		for _, direction := range Directions {
			if fish.Typ.Value() >= block.Metas[direction].Length {
				return true
			}
		}
	}
	return false
}

// UsageDirections 可用的方位
func UsageDirections(blocks []Block, fish Fish) []Direction {
	directions := make([]Direction, 0, 2)
	for _, block := range blocks {
		for _, direction := range Directions {
			if fish.Typ.Value() >= block.Metas[direction].Length {
				directions = append(directions, direction)
			}
		}
	}
	return directions
}

//+----------------------------------------------------------------+
//+					         Fish Handle	                       +
//+----------------------------------------------------------------+

// FishType 鱼的种类
type FishType uint8

func (f FishType) None() bool {
	return f == None
}

func (f FishType) IsFish() bool {
	return f > 0 && f < 5
}

func (f FishType) Value() int {
	return int(f)
}

const (
	None      FishType = iota // 什么也没有，长度0
	Blowfish                  // 河豚，长度1
	Crabs                     // 螃蟹，长度2
	Shark                     // 鲨鱼，长度3
	BlueWhale                 // 蓝鲸，长度4

	Blank FishType = 20 // 空格子，长度1
)

const (
	JackpotGrand FishType = 21 + iota // Grand Jackpot，长度1
	JackpotMajor
	JackpotMinor
	JackpotMini
)

// Fish 鱼
type Fish struct {
	State      int      // 0：未使用，1：已使用
	Index      int      // 所属鱼池位置
	Typ        FishType // 鱼的种类
	Seq        int      // 出现的顺序
	Horizontal bool     // 是否水平摆放
}

// Fishpond 鱼池
type Fishpond struct {
	ShowNum int    // 已经展示的鱼的计数
	Fishes  []Fish // 所有鱼
}

func NewFishpond(fishes []Fish) *Fishpond {
	fishpond := &Fishpond{
		Fishes: fishes,
	}
	fishpond.Tidy()
	return fishpond
}

// Tidy 整理索引
func (f *Fishpond) Tidy() *Fishpond {
	if len(f.Fishes) == 0 {
		return nil
	}
	for i := range f.Fishes {
		f.Fishes[i].Index = i
	}
	sort.Slice(f.Fishes, func(i, j int) bool {
		return f.Fishes[i].Typ >= f.Fishes[j].Typ
	})
	return f
}

// Get 获取一条未被放置的鱼
func (f *Fishpond) Get() (*Fish, error) {
	if len(f.Fishes) == 0 {
		return nil, ErrNoMoreFish
	}
	use := make([]Fish, 0, len(f.Fishes)) // 可以使用的鱼
	for _, fish := range f.Fishes {
		if fish.State == 0 {
			use = append(use, fish)
		}
	}
	if len(use) == 0 {
		return nil, ErrNoMoreFish
	}
	if len(use) == 1 {
		return &use[0], nil
	}
	fish := use[rand.Intn(len(use))]
	return &fish, nil
}

// NeedToPlace 需要放置的鱼
func (f *Fishpond) NeedToPlace() []Fish {
	use := make([]Fish, 0, len(f.Fishes)) // 可以使用的鱼
	for _, fish := range f.Fishes {
		if fish.State == 0 {
			use = append(use, fish)
		}
	}
	return use
}

// UpdateState 更新鱼的状态为已使用
func (f *Fishpond) UpdateState(index int, horizontal bool) {
	if len(f.Fishes) == 0 {
		return
	}
	f.ShowNum++
	f.Fishes[index].State = 1
	f.Fishes[index].Seq = f.ShowNum
	f.Fishes[index].Horizontal = horizontal
}

// Copy 获得副本
func (f *Fishpond) Copy() Fishpond {
	fishes := make([]Fish, len(f.Fishes))
	copy(fishes, f.Fishes)
	return Fishpond{
		Fishes: fishes,
	}
}

// Reverse 反转数组
func Reverse(blocks []Block) []Block {
	result := make([]Block, len(blocks))
	if len(blocks) == 0 {
		return result
	}
	for i, block := range blocks {
		result[len(blocks)-i-1] = block
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
