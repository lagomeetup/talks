package main

import (
	"fmt"
	"os/user"
)

func main() {
	me, err := user.Current()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Username: %s\n", me.Username)
		fmt.Printf("Name: %s\n", me.Name)
		fmt.Printf("HomeDir: %s\n", me.HomeDir)
	}
}
