package synctest

import "testing"

func TestMutex(t *testing.T) {
	Mutex()
}

func TestRWMutex(t *testing.T) {
	RWMutex()
}

func TestWaitGroup(t *testing.T) {
	WaitGroup()
}

func TestCond(t *testing.T) {
	Cond()
}

func TestAtomic(t *testing.T) {
	Atomic()
}

func TestOnce(t *testing.T) {
	Once()
}

func TestMap(t *testing.T) {
	Map()
}
