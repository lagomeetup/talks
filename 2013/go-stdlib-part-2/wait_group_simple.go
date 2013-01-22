package main

import "fmt"
import "sync"

var V = []string { "a", "b", "c", "d" }

func main() {
	var wg sync.WaitGroup
	for _, s := range V {
		wg.Add(1)
		go func() {
			fmt.Println(s)
			wg.Done()
		}()
	}
	done.Wait()
}
