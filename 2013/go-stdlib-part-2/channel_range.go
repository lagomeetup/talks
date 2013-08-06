package main

import (
	"fmt"
)

func main() {

	blocking := make(chan int, 5)

	for i := 0; i < 10; i++ {
		blocking <- i
	}

	//close(blocking)

	for value := range blocking {
		fmt.Printf("Completed with %v\n", value)
	}

}
