package checksvn

import (
	"context"
	"errors"
	"strings"
	"time"
)

var (
	ErrEmptyMsg         = errors.New("获取SVN消息为空")
	ErrEmptyRows        = errors.New("SVN空行错误")
	CheckInterval int64 = 300 // 检测间隔，和定时任务一致
)

// RunSvnCmd 运行SVN命令行
func RunSvnCmd(cmd Cmd, auth *SVNConfig, s Server) (*Row, error) {
	output, err := Command(``, formatCmd(cmd, auth, s))
	if err != nil {
		return nil, err
	}
	output = strings.TrimSpace(strings.ReplaceAll(output, splitLine, ""))
	StdOut([]byte(output), `RunSvnCmd`)
	rows, step := SvnCmdMap[cmd](&output)
	if step == Interrupt {
		return &rows[0], nil
	}
	// TODO after process
	return &rows[0], nil
}

// TickCall 定时器调用
func TickCall(ctx context.Context, ch <-chan time.Time, f func()) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-ch:
			f()
		}
	}
}
