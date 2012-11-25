package write

import (
        "encoding/binary"
        "io/ioutil"
        "testing"
)

type Action struct {
        U1 uint64
        U2 uint64
        U3 uint64
}

const fileSize = 1 << 16

func BenchmarkBinaryWrite64k(b *testing.B) {
        var action Action
        b.SetBytes(fileSize)
        for i := 0; i < b.N; i++ {
                for totalSize := 0; totalSize < fileSize; {
                        binary.Write(ioutil.Discard, binary.LittleEndian, &action)
                        totalSize += 24
                }
        }
}