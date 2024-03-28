package casino

import (
	"fmt"
	"math/rand"
	"time"
)

// Inputer 输入
type Inputer interface {
	Value() interface{}
}

// Outer 输出
type Outer interface {
	Default(error) Outer
	Err() error
	Value() interface{} // TODO implement interface{ func... }
}

// GameOperator 游戏玩法操作接口
type GameOperator interface {
	Do(Inputer) Outer // 执行操作
}

type User struct {
	UID      int32 `json:"UID"`      // 用户的ID
	RoundNum int32 `json:"RoundNum"` // 当前轮次进度
	Game     *Game // 当前游戏进度
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewUser(uid int32) *User {
	return &User{UID: uid}
}

func (user *User) Init() error {
	user.RoundNum = 1
	user.Game = NewGame()
	return nil
}

func (user *User) Print() string {
	return fmt.Sprintf("UID:%d\n%s", user.UID, user.Game.Print())
}
