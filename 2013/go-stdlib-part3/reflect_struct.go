package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := struct {
		Something int `tag:"something"`
		string  // Anonymous field
	}{0,"X"}

	a.Something = 0

	pt := reflect.TypeOf(a)

	for i := 0; i < pt.NumField(); i++ {
		fdesc := pt.Field(i)
		fmt.Printf("\n");
		fmt.Printf("\tName:%v\n", fdesc.Name)
		fmt.Printf("\tPkgPath:%v\n", fdesc.PkgPath)
		fmt.Printf("\tType:%v\n", fdesc.Type)
		fmt.Printf("\tTag:%v\n", fdesc.Tag)
		fmt.Printf("\t\t:%v\n",fdesc.Tag.Get("tag"))
		fmt.Printf("\tOffset:%v\n", fdesc.Offset)
		fmt.Printf("\tIndex:%v\n", fdesc.Offset)
		fmt.Printf("\tAnonymous:%v\n", fdesc.Anonymous)
    }

}
