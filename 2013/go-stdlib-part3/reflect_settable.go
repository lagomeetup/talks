package main

import (
	"fmt"
	"reflect"
)

func main() {
	// START1 OMIT
	var value int = 10
	var V reflect.Value = reflect.ValueOf(value)
	fmt.Printf("CanSet:%v\n", V.CanSet())
	// END1 OMIT

	// START2 OMIT
	V = reflect.ValueOf(&value)
	fmt.Printf("CanSet:%v\n", V.CanSet())
	// END2 OMIT

	// START3 OMIT
	V = reflect.ValueOf(&value).Elem()
	fmt.Printf("CanSet:%v\n", V.CanSet())

	V.SetInt(20)
	fmt.Printf("NewValue: %d\n", value)
	// END3 OMIT

	var arrayValue []int = []int{1, 2}
	V = reflect.ValueOf(&arrayValue)
	fmt.Printf("CanSet:%v\n", V.CanSet())

	V = reflect.ValueOf(&arrayValue)
	V = reflect.Indirect(V)
	fmt.Printf("CanSet:%v\n", V.CanSet())

}
