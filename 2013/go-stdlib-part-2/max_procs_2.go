package main

import (
	"fmt"
)

// START1 OMIT
func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}
// END1 OMIT

func main() {
	fmt.Printf ("Parallel unit: %d\n",MaxParallelism())
}
