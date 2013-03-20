package main

// #include <string.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	str := "Gophers Rock!"
	dup := C.strdup(C.CString(str))
	defer func() {
		C.free(unsafe.Pointer(dup))
	}()

	gostr := C.GoString(dup)
	fmt.Println(gostr)
}
