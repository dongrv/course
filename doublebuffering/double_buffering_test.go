package doublebuffering

import (
	"testing"
)

func TestExchange_RWGoroutine(t *testing.T) {
	exchange := NewExchange()
	exchange.RWGoroutine()
}
