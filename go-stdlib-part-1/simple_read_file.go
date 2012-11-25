package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// <exists OMIT
// PathExists returns whether the given path exists or not
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path) // HL
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	} // HL
	return false, err
}

// exists> OMIT

func main() {
	// {path OMIT
	file_path := "$HOME/.gitconfig"
	file_path_clean := filepath.Clean(os.ExpandEnv(file_path)) // HL
	if exists, path_err := PathExists(file_path_clean); !exists {
		// path} OMIT
		if path_err != nil {
			fmt.Printf("File not found %s (%s) \n", file_path_clean, path_err)
		} else {
			fmt.Printf("File not found %s\n", file_path_clean)
		}
		os.Exit(-1)
	}

	// START1 OMIT
	fd, err := os.OpenFile(file_path_clean, os.O_RDONLY, 0) // HL
	if err != nil {
		os.Exit(-1)
	}
	defer fd.Close() // HL
	// STOP1 OMIT
	// START2 OMIT
	buf := make([]byte, os.Getpagesize()) // HL
	count, read_err := fd.Read(buf)       // HL
	sum := 0
	for read_err == nil && count > 0 {
		count, read_err = fd.Read(buf) // HL
		sum += count
	}
	// STOP2 OMIT

	switch read_err {
	case nil:
		// do nothing
	case io.EOF:
		fmt.Printf("%s Total Size: %d\n", file_path_clean, sum)
	default:
		os.Exit(-1)
	}
}
