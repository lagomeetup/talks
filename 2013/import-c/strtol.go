package main

// #include <stdlib.h>
import "C"

import (
	"fmt"
)

func main() {
	// Notice you can pass nil as NULL
	i := C.strtol(C.CString("10"), nil, 10)
	fmt.Println(i)

	// strtol will set errno on errors (in this case, base too large)
	i, err := C.strtol(C.CString("10"), nil, 100)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
}
