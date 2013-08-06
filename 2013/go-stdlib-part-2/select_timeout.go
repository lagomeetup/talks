package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	handle := func() {
		time.Sleep(time.Second * 2)
		c <- 1
	}
	go handle()

	select {
	case m := <-c:
		fmt.Printf("ok: %d\n", m)
	case <-time.After(500 * time.Second):
		fmt.Println("timed out")
	default:
		fmt.Println(".")
	}
}
