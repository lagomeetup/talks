package main

import (
	"fmt"
	"reflect"
)

// from http://stackoverflow.com/questions/6395076/in-golang-using-reflect-how-do-you-set-the-value-of-a-struct-field

func main() {
	// START1 OMIT
	type t struct {
		N int
	}
	var n = t{42}
	fmt.Println(n.N)
	ps := reflect.ValueOf(&n)
	s := ps.Elem()
	if s.Kind() == reflect.Struct {
		f := s.FieldByName("N")
		if f.IsValid() {
			if f.CanSet() {
				if f.Kind() == reflect.Int {
					x := int64(7)
					if !f.OverflowInt(x) {
						f.SetInt(x)
					}
				}
			}
		}
	}
	// END1 OMIT

	// N at end
	fmt.Println(n.N)
}
