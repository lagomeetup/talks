package main

import (
    "bytes"
    "encoding/binary"
    "fmt"
)

func main() {
    // START OMIT
    var val struct {
        A, B    uint8
        C   [2]uint8
        D   float32
    }
    b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
    buf := bytes.NewBuffer(b)
    err := binary.Read(buf, binary.LittleEndian, &val)
    if err != nil {
        fmt.Println("binary.Read failed:", err)
    }
    fmt.Printf("%#v", val)
    // END OMIT
}