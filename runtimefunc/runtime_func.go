package runtimefunc

import "runtime"

func RuntimeFuncs() {
	runtime.GOMAXPROCS(1)
	runtime.GC()
	runtime.Version()
	runtime.Gosched()
	runtime.NumCPU()
	runtime.NumGoroutine()
	runtime.Goexit()
}
