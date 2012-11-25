package main

import (
    "os"
    "bufio"
)

func main() {
    // START1 OMIT
    f := os.Stdout
    inner := bufio.NewWriter(f)
    w := bufio.NewWriterSize(inner, 8192)
    for i:=0; i < 10; i++ {
        w.WriteString("ABC\r\n")
    }
    w.Flush()
    f.Close()
    // START2 OMIT
}