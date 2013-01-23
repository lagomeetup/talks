package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	for i := 0; i < 100; i++ {
		go func() {
			for {
			}
		}()
	}
	var end int
	fmt.Scanf("Wait For It:%d", &end)
	// END1 OMIT
}
