package main

import (
	"fmt"
	"reflect"
)

func meta(a interface{}) {
	pt := reflect.TypeOf(a)
	pv := reflect.ValueOf(a)

	fmt.Printf("\n")
	for i := 0; i < pt.Len(); i++ {
		elementValue := pv.Index(i)
		fmt.Printf("%v:%v:%v:%v\n", i, elementValue.Interface(), elementValue.Type(), elementValue.Kind())
    }
    fmt.Printf("\n")
}

func main() {
	a := [10]int{1,2,3,4,5,6,7,8,9,10}

	meta(a)
   
    b := [3]interface{}{1,"IamString", struct { int }{0} }

    meta(b)
}
