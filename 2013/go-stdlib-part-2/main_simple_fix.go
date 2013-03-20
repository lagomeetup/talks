package main

import (
	"fmt"
)

func status() {
	fmt.Printf("It works !")
}


func main() {
	go status()
	for {}
}