package main

import (
	"fmt"
	"template/action"
	"template/bingo"
)

func main() {
	user := casino.NewUser(10001)
	if err := user.Init(); err != nil {
		return
	}

	_ = action.PlayAction(user, &action.PlayReq{X: 0, Y: 1})
	_ = action.PlayAction(user, &action.PlayReq{X: 1, Y: 1})
	err := action.PlayAction(user, &action.PlayReq{X: 2, Y: 1})
	fmt.Printf("游戏进度:%s 错误:%v\n", user.Print(), err)

}
