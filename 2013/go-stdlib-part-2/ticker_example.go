package main

import (
	"time"
	"fmt"
)

func main() {
	for _ = range time.Tick( time.Millisecond * 500 ) {
		fmt.Printf("Tick ...")
	}
}
