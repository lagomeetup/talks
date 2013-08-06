package main

import (
	"fmt"
)

type Callback func()

// START2 OMIT
func mkCallback(i int) Callback {
	return func() { fmt.Printf("%d\n", i) } // HL
}

// END2 OMIT

func main() {
	// START1 OMIT
	funcArray := make([]Callback, 0)

	for i := 0; i < 10; i++ {
		aNewFunc := mkCallback(i) // HL
		funcArray = append(funcArray, aNewFunc)
	}

	for i := 0; i < 10; i++ {
		funcArray[i]()
	}
	// END1 OMIT
}
