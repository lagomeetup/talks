package main

import (
	"runtime"
)

func main() {
	ch := make(chan int)
	ch2 := make(chan int)
	for {
		go func() {
			ch <- 0
		}()
		runtime.Gosched()
	}
	ch2 <- 0
}
