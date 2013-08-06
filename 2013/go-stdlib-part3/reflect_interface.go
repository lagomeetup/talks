package main

import (
	"fmt"
	"reflect"
)

// START3 OMIT
type SampleInterface interface {
	PowerReflection()
}

type Implementation struct{}

func (*Implementation) PowerReflection() {}

// END3 OMIT

// START1 OMIT
func meta(pv reflect.Type) {
	fmt.Printf("\n%v\n", pv)

	for i := 0; i < pv.NumMethod(); i++ {
		methodValue := pv.Method(i)
		fmt.Printf("\tName:%v\n", methodValue.Name)
		fmt.Printf("\tPkgPath:%v\n", methodValue.PkgPath)
		fmt.Printf("\tType:%v\n", methodValue.Type)
		fmt.Printf("\tFunc:%v\n", methodValue.Func)
		fmt.Printf("\tIndex:%v\n", methodValue.Index)
	}
}

// END1 OMIT

func main() {
	// START2 OMIT
	pv := reflect.TypeOf((*SampleInterface)(nil)).Elem() // HL
	meta(pv)
	a := &Implementation{}
	meta(reflect.TypeOf(a))
	// END2 OMIT
}
