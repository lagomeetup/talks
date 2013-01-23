package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	blocking := make(chan int, 1)
	// blocking <-1
	close(blocking)
	if value, ok := <-blocking; ok {
	   fmt.Printf("Completed with %v\n", value)
	} else {
	   fmt.Printf("Channel is closed\n")
	}
	// END1 OMIT
}
