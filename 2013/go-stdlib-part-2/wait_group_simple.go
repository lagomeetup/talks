package main

import "fmt"
import "sync"

func main() {
	// START1 OMIT
	var V = []string{"a", "b", "c", "d"}
	var wg sync.WaitGroup // HL
	for _, s := range V {
		wg.Add(1) // HL
		go func() {
			fmt.Println(s)
			wg.Done()  // HL
		}()
	}
	wg.Wait()  // HL
	// END1 OMIT
}
