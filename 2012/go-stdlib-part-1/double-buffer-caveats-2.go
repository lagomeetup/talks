package main

import (
    "os"
    "bufio"
)

func main() {
    // START1 OMIT
    f := os.Stdout
    innerBuffer := bufio.NewWriter(f) // Default buffer size is 4096 bytes
    w:= bufio.NewWriterSize(innerBuffer, 8192)
    for i:=0; i < 10; i++ {  
        w.WriteString("ABC\r\n")
    }
    w.Flush()   // HL
    innerBuffer.Flush()  // HL
    f.Close()
    // START2 OMIT
}