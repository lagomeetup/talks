package main

import (
	"crypto/md5"
	"fmt"
	"flag"
	"io"
	"os"
)

// open open a file, the filename is - will return os.Stdin
func open(filename string) (io.ReadCloser, error) {
	if filename == "-" {
		return os.Stdin, nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return file, err
}

func main() {
	var filename string
	flag.StringVar(&filename, "file", "-", "file name")
	flag.Parse()

	file, err := open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: can't open %s - %s\n", filename, err)
		os.Exit(1)
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		fmt.Fprintf(os.Stderr, "error: can't read %s - %s\n", filename, err)
		os.Exit(1)
	}

	fmt.Printf("%x\n", hash.Sum(nil))
}
