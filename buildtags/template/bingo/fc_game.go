//go:build fatcat
// +build fatcat

package casino

import (
	"errors"
	"fmt"
	"sync"
)

const (

	// 棋盘大小

	BoardX = 3
	BoardY = 3

	HitBonus = 1 // 命中奖励
)

var (
	ErrAssertFail   = errors.New("assert pos data fail")
	ErrHitNothing   = errors.New("hit nothing")
	ErrInvalidClick = errors.New("invalid click")
)

// CustomConfig 棋盘和奖励配置
type CustomConfig struct {
	Board   [BoardX][BoardY]int32 // 棋盘：3x3的二维数组
	Weights []int32               // 筛选奖励的权重
	Awards  []int32               // 命中时奖励的积分
}

func (c *CustomConfig) Copy() CustomConfig {
	newWeight := make([]int32, len(config.Custom.Weights))
	newAwards := make([]int32, len(config.Custom.Awards))
	copy(newWeight, config.Custom.Weights)
	copy(newAwards, config.Custom.Awards)
	return CustomConfig{
		Board:   config.Custom.Board,
		Weights: newWeight,
		Awards:  newAwards,
	}
}

type InputVar struct {
	X, Y int
}

func NewInputVar() InputVar {
	return InputVar{}
}

func (input *InputVar) Value() interface{} {
	return *input
}

func (input *InputVar) set(x, y int) InputVar {
	return InputVar{x, y}
}

func (input *InputVar) verify() bool {
	return (0 <= input.X && input.X < BoardX) && (0 <= input.Y && input.Y < BoardY)
}

type Result struct {
	err     error
	Payload interface{} // 信息载体
}

func (result Result) Default(err error) Outer {
	return Result{err: err}
}

func (result Result) Err() error {
	return result.err
}

func (result Result) Value() interface{} {
	return result.Payload
}

// Game 自定义游戏数据载体
type Game struct {
	mu        sync.RWMutex
	PickNum   int32                 `json:"PickNum"`   // 剩余pick次数
	Bonus     int32                 `json:"Bonus"`     // 奖励的积分
	PickedPos map[InputVar]struct{} `json:"PickedPos"` // 点击过的位置
	Done      bool                  // 是否完成游戏：找到棋盘上所有为1的位置
}

func NewGame() *Game {
	return &Game{
		PickedPos: make(map[InputVar]struct{}),
		PickNum:   10,
	}
}

// CanClick 是否可以点击目标位置
func (game *Game) CanClick(pos InputVar) bool {
	if game.PickNum == 0 {
		return false
	}
	game.mu.RLock()
	defer game.mu.RUnlock()
	_, ok := game.PickedPos[pos]
	return !ok
}

// 占位
func (game *Game) setPos(p InputVar) {
	game.mu.Lock()
	defer game.mu.Unlock()
	game.PickedPos[p] = struct{}{}
}

// Do 执行核心逻辑
func (game *Game) Do(input Inputer) Outer {
	result := Result{}

	// 检查参数
	pos, ok := input.Value().(InputVar)
	if !ok {
		return result.Default(ErrAssertFail)
	}
	if !pos.verify() {
		return result.Default(ErrInvalidClick)
	}
	if !game.CanClick(pos) {
		return result.Default(ErrInvalidClick)
	}
	if config.Custom.Board[pos.X][pos.Y] != HitBonus {
		return result.Default(ErrHitNothing) // 此点击未中奖
	}

	// 计算奖励内容
	copies := config.Custom.Copy()
	index := WeightChosen(copies.Weights)
	if index < 0 {
		return result.Default(errors.New("except")) // 配置正常可以排除此报错
	}

	// 发放奖励
	bonus := copies.Awards[index]
	game.PickNum--
	game.Bonus += bonus
	game.PickedPos[pos] = struct{}{}

	return result.Default(nil) // 没返回说明中奖了
}

func (game *Game) Print() string {
	var coordinates string
	for inputVar := range game.PickedPos {
		coordinates += fmt.Sprintf("(%d, %d)\n", inputVar.X, inputVar.Y)
	}
	return fmt.Sprintf("PickNum:%d\nBonus:%d\n点击过的坐标:\n%s\n", game.PickNum, game.Bonus, coordinates)
}
