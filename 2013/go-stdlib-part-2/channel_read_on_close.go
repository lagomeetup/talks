package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	blocking := make(chan int)
	// blocking := make(chan string)
	close(blocking)
	value := <-blocking
	fmt.Printf("Completed with %v\n", value)
	// END1 OMIT
}
