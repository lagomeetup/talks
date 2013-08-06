package main

import (
	"fmt"
)

type MyType struct {
	state int
}

func main() {
     // START1 OMIT
	my_channel := make(chan MyType, 1)
	reference := MyType{}
	reference.state = 100 // HL
	my_channel <- reference
	reference.state = 200 // HL
	fmt.Printf("%v\n", reference)
	received := <-my_channel
	fmt.Printf("%v\n", received)
	// END1 OMIT
}
