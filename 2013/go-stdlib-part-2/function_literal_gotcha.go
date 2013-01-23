package main

import (
	"fmt"
)

type Callback func()

func main() {
	// START1 OMIT
	funcArray := make([]Callback, 0)

	for i := 0; i < 10; i++ {
		aNewFunc := func() { fmt.Printf("%d\n", i) } // HL
		funcArray = append(funcArray, aNewFunc)
	}

	for i := 0; i < 10; i++ {
		funcArray[i]()
	}
	// END1 OMIT
}
