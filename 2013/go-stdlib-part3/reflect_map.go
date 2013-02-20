package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := map[string]interface{}{ 
		"a" : nil,
		"b" : 10,
 	}

 	pt := reflect.TypeOf(a)
 	pv := reflect.ValueOf(a)

    for _, k := range pv.MapKeys() {
    	fmt.Printf("\n")
		fmt.Printf("\tName:%v\n", k.Name)
		fmt.Printf("\tPkgPath:%v\n", fdesc.PkgPath)
		fmt.Printf("\tType:%v\n", fdesc.Type)
		fmt.Printf("\tTag:%v\n", fdesc.Tag)
		fmt.Printf("\t\t:%v\n",fdesc.Tag.Get("tag"))
		fmt.Printf("\tOffset:%v\n", fdesc.Offset)
		fmt.Printf("\tIndex:%v\n", fdesc.Offset)
		fmt.Printf("\tAnonymous:%v\n", fdesc.Anonymous)
    }

};
