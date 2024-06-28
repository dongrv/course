package printmem

import (
	"fmt"
	"runtime"
)

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	tomb := func(b uint64) uint64 {
		return b / 1024 / 1024
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", tomb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", tomb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", tomb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
