package checksvn

import (
	"context"
	"testing"
	"time"
)

var (
	config = NewConfig(`./Config.json`)
)

func TestRunSvnCmd(t *testing.T) {
	updateLog, err := RunSvnCmd(ShowLog, config.Svn, Business) // 获取更新日志
	if err != nil {
		t.Fatal(err)
	}
	if Empty(*updateLog) {
		return // 空行不处理
	}
	_, err = RunSvnCmd(Checkout, config.Svn, Business) // checkout失败
	if err != nil {
		t.Fatal(err)
	}

	output, err := Command(``, config.GoBuildScript)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("output: %v", output)
	err = NotifyGroup(WrapMessage(*updateLog, true, ``), config.Msg)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTickCall(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ticker := time.NewTicker(5 * time.Minute)

	defer func() {
		cancel()
		ticker.Stop()
	}()

	TickCall(ctx, ticker.C, func() {
		row, err := RunSvnCmd(ShowLog, config.Svn, Business)
		if err != nil {
			t.Fatal(err)
		}
		err = NotifyGroup(WrapMessage(*row, true, ``), config.Msg)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(row)
	})

}
