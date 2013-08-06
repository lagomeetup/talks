package main

import (
	"fmt"
	"runtime"
)

// START1 OMIT
func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs > numCPU {
		return maxProcs
	}
	return numCPU
}
// END1 OMIT

func main() {
	runtime.GOMAXPROCS(MaxParallelism())
	fmt.Printf ("Parallel unit: %d\n",MaxParallelism())
}
