package checksvn

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
)

var ErrEmptyMsg = errors.New("获取SVN消息为空")

var CheckInterval int64 = 5 * 60 // 检测间隔，和定时任务一致

// RunSvnCmd 运行SVN命令行
func RunSvnCmd(svnCmd Cmd, auth *SvnAuth, bot *ChatGroupBot) error {
	cmd := fmt.Sprintf(svnCmd.String(), auth.Host, auth.Username, auth.Password)
	output, err := Command(cmd)
	if err != nil {
		return err
	}
	output = strings.TrimSpace(strings.ReplaceAll(output, splitLine, ""))
	results := SvnCMDMap[ShowLog](&output)
	if len(results) == 0 {
		return ErrEmptyMsg
	}
	StdOut([]byte(output), `RunSvnCmd`)
	row := results[0] // 取最近行
	if time.Now().Unix()-row.Timestamp <= CheckInterval {
		err = NotifyGroup(WrapMessage(row), bot)
	}
	return err
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
