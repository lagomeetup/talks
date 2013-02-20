package main

import (
	"fmt"
	"reflect"
)

func main() {
	// START1 OMIT
	type MyInt int

	var i int
	var j MyInt

	T := reflect.TypeOf(i)
	//V := reflect.ValueOf(i)
	fmt.Printf("Type i: %v\n", T)
	fmt.Printf("Kind i: %v\n", T.Kind())

	T1 := reflect.TypeOf(j)
	//V1 := reflect.ValueOf(j)
	fmt.Printf("Type j: %v\n", T1)
	fmt.Printf("Kind j: %v\n", T1.Kind())

	var jj int = int(j) 

	T2 := reflect.TypeOf(jj)
	fmt.Printf("Type i: %v\n", T2)
	fmt.Printf("Kind i: %v\n", T2.Kind())
	// END1 OMIT
}