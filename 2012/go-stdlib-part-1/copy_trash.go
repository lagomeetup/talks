package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
)

func main() {
// START OMIT
	file, err := ioutil.ReadFile(os.ExpandEnv("$HOME/.gitconfig"))
	if err != nil {
		println(err)
	}
	io.Copy(ioutil.Discard, bytes.NewReader(file))
} // STOP OMIT
