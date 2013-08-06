package main

import (
	"reflect"
)

// START1 OMIT
type SampleInterface interface  {
	PowerReflection()
}
// END1 OMIT

func main() {
	// START2 OMIT
	pv := reflect.TypeOf(SampleInterface)
	// END2 OMIT
}
