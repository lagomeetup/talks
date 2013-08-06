package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myValue int 
	myValue = 42
	T  := reflect.TypeOf(myValue)
	V  := reflect.ValueOf(myValue)
	fmt.Printf("Type: %v\n",T)
	fmt.Printf("Value: %v\n", V)
	fmt.Printf("Value.Interface: %v\n", V.Interface())
	fmt.Printf("Settable: %v\n", V.CanSet())
	switch V.Kind() {
	case reflect.Ptr,reflect.Interface:
		fmt.Printf("Elem: %v\n", V.Elem())
	}

	fmt.Printf("Zero: %v\n", reflect.Zero(T).Interface())

	var myPointerToValue *int
	myPointerToValue = &myValue
	fmt.Printf("Type: %v\n",reflect.TypeOf(myPointerToValue))
	fmt.Printf("Value: %v\n",reflect.ValueOf(myPointerToValue))
	fmt.Printf("Value.Interface: %v\n",reflect.ValueOf(myPointerToValue).Interface())

}
