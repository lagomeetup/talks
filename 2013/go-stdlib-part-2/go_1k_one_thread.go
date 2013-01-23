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

	for {}
	// END1 OMIT
}
