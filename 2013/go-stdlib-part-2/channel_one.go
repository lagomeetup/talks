package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	blocking := make(chan int, 1)
	go func () { blocking <- 1 }()
	// go func () { blocking <- 1 }()
	<-blocking
	<-blocking
	fmt.Printf("Completed\n")
	// END1 OMIT
}
