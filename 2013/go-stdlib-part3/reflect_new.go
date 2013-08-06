package main

import (
	"fmt"
	"reflect"
)

func main() {
	// START1 OMIT
	type SomeStruct struct {
		aa [10]int
		bb [20]int
	}

	var intPtr SomeStruct
	intPtr2 := reflect.New(reflect.TypeOf(intPtr))

	fmt.Printf("New (Pointer to T): %v: Type %v\n", intPtr2.Elem().Interface(), intPtr2.Type())

	intPtr3 := reflect.Zero(reflect.TypeOf(intPtr))

	fmt.Printf("Zero (T): %v: Type:%v\n", intPtr3.Interface(), intPtr3.Type())
	// END1 OMIT
}
