package main

// meanwhile at stdlib.h:
// typedef struct
//   {
//     int quot;			/* Quotient.  */
//     int rem;			/* Remainder.  */
//   } div_t;

// #include <stdlib.h>
import "C"

import (
	"fmt"
)

func main() {
	var dt C.div_t

	dt = C.div(16, 6)
	fmt.Printf("quot: %d, rem: %d\n", dt.quot, dt.rem)
}

