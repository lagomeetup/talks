package main

import "fmt"
import "sync"

func main() {
	// START1  OMIT
	var V = []string{"a", "b", "c", "d"}
	var wg sync.WaitGroup
	for _, s := range V {
		wg.Add(1)
		go func(s string) { 
			defer wg.Done() // HL
			fmt.Printf(s)
		}(s) 
	}
	wg.Wait()
	// STOP1  OMIT
}
