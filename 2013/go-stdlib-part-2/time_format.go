package main

import (
	"fmt"
	"time"
	"github.com/tebeka/strftime"
)

func main() {
	fmt.Println(time.Now().Format("2006 Jan 01 02 15 04 05 MST -0700"))
	fmt.Println(strftime.Format("%Y %b %m %d %H %M %S %Z", time.Now()))
}
