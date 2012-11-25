package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
    // START OMIT
	reader := bufio.NewReader(os.Stdin) // HL
	var err error = nil
	var result string

	for err == nil {
		result, err = reader.ReadString('\n') // HL
		println(result)
	}

	if err != nil && err != io.EOF {
		println(err.Error())
	}
	// STOP OMIT
}
