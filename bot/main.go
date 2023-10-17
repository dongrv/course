package main

import (
	"bot/dialog"
	"bot/msg"
	"bot/msg/protocol"
	"bot/robot"
	"context"
)

var protoFile = "./ProtoConfig.json" // TODO 改成本地文件路径

func main() {

	dialog.InitLogger()          // 设定日志格式
	protocol.Register(protoFile) // 注册所有协议
	msg.RegisterFunc()           // 注册处理函数

	center := robot.NewCenter(2, `robot`, `:2001`)

	ctx, cancel := context.WithCancel(context.Background())
	center.Activate(ctx) // 激活机器人
	dialog.WaitQuit()    // 监听退出信号
	cancel()             // 取消上下文
	center.Wait()        // 等待所有协程退出

}
