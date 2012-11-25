package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
)

func check(e error) {
	if e != nil && e != io.EOF {
		log.Fatal(e)
	}
}
func main() {
	//WRITE1 OMIT
	txt := "someverylonglineoftext to be compressed\n"
	var buf bytes.Buffer
	bufzip := gzip.NewWriter(&buf) // writes to buf

	bufWrite := bufio.NewWriter(bufzip) // bufWrite writes to bufzip
	bufWrite.WriteString(txt)
	bufWrite.Flush()

	bufzip.Close()  

	//WRITE2 OMIT

	var (
		part string
	)

	// READ1 OMIT
	zipread, err := gzip.NewReader(&buf)
	check(err)
	reader := bufio.NewReader(zipread)

	for ; err == nil; part, err = reader.ReadString('\n') {
		fmt.Print(part)
	}
	if err == io.EOF {
		err = nil
	}
	// READ2 OMIT
	fmt.Println("Complete")
}
