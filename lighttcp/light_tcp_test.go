package lighttcp

import (
	"context"
	"testing"
)

func TestGuard_StartRW(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g := &Guard{}
	g.Init()
	g.OnSocket(ctx, TCPConfig{`tcp`, `:2001`})

}
