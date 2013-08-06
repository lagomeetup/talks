package main

import (
	"fmt"
	"reflect"
)

func main() {
	// START1 OMIT
	a := map[string]interface{}{
		"a": nil,
		"b": 10,
	}

	pv := reflect.ValueOf(a)

	fmt.Printf("\n%v%n", a)
	for _, k := range pv.MapKeys() {
		elementValue := pv.MapIndex(k)
		fmt.Printf("\tKey:%v\n", k.String())
		fmt.Printf("\tValue:%v:%v:%v\n", elementValue.Interface(), elementValue.Type(), elementValue.Kind())
	}
	// END1 OMIT
}
