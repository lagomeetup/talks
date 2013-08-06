package main

import (
	"fmt"
	"runtime"
)

func main() {
	// START1 OMIT
	for i := 0; i < 1000; i++ {
		go func() {
			for {
			}
		}()
		fmt.Printf("GoRoutines: %d \n", runtime.NumGoroutine())
	}

	for {
		fmt.Printf("GoRoutines: %d \n", runtime.NumGoroutine())
	}
	// END1 OMIT
}
