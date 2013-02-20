package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	sl := a[0:5]

	ra := reflect.Indirect(reflect.ValueOf(&a)).Slice(0, 5)

	fmt.Printf("%v\n%v\n", sl, ra.Interface())
}
