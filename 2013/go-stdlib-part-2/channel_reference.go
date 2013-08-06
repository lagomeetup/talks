package main

import (
	"fmt"
)

// START1 OMIT
type MyType struct {
	state int
}

func main() {
	my_channel := make(chan *MyType, 1)
	reference := new(MyType)
	reference.state = 100 // HL
	my_channel <- reference
	reference.state = 200 // HL
	fmt.Printf("%v\n", reference)
	received := <-my_channel
	fmt.Printf("%v\n", received)
	// END1 OMIT
}
