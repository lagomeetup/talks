package main

import (
	"fmt"
	"reflect"
)

// START1 OMIT
type Person struct{}

func (p *Person) Run() {
	fmt.Println("person running")
}

func doit(object interface{}, method interface{}) {
	v := reflect.ValueOf(object)
	f := reflect.ValueOf(method)
	f.Call([]reflect.Value{v})
}

func main() {
	p := new(Person)
	doit(p, (*Person).Run)
	// END1 OMIT
}
