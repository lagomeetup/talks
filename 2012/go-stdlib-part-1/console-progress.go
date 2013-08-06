
package main

import (
    "os"
    "bufio"
    "time"
)       

func main(){
    // START OMIT
    w := bufio.NewWriter(os.Stdout)
    c := time.Tick(1 * time.Second)

    for _ = range c {
        w.WriteString("tick")
        //w.Flush()
    }
    // STOP OMIT
} 