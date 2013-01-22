package main

import (
	"fmt"
)

// START1 OMIT
type Callback func()  // HL
// END1 OMIT

func main() {
	//START2 OMIT
	var my_function Callback  // HL
	my_function = func() {
		fmt.Printf("%d\n", 10)
	}
	my_function()
	//END2 OMIT
}
