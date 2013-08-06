package main

import (
	"fmt"
	"reflect"
)

func main() {
	// START1 OMIT
	var i int = 10
	var a chan int

	channel := reflect.MakeChan(reflect.TypeOf(a), 1)

	channel.Send(reflect.ValueOf(i))
	value, ok := channel.Recv()

	if ok {
		fmt.Printf("Magic! %v\n", value.Interface())
	} else {
		fmt.Printf("Something went wrong :(")
	}
	// END1 OMIT
}
