package main

import "fmt"
import "sync"

func main() {
	// START1  OMIT
	var V = []string{"a", "b", "c", "d"}
	var wg sync.WaitGroup
	for _, s := range V {
		wg.Add(1)
		go func(s string) { // HL
			fmt.Println(s)
			wg.Done()
		}(s) // HL
	}
	wg.Wait()
	// STOP1  OMIT
}
