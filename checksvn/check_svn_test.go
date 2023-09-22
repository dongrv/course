package checksvn

import (
	"context"
	"testing"
	"time"
)

var (
	auth = NewSvnAuth(``, ``, ``)
	bot  = NewChatGroupBot(``)
)

func TestRunSvnCmd(t *testing.T) {
	t.Log(RunSvnCmd(ShowLog, auth, bot))
}

func TestTickCall(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	ticker := time.NewTicker(5 * time.Minute)

	defer func() {
		cancel()
		ticker.Stop()
	}()

	TickCall(ctx, ticker.C, func() {
		_ = RunSvnCmd(ShowLog, auth, bot)
	})

}
