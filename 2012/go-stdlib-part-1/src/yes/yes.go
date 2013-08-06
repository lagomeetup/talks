// Ignore this this is documentation 
// aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaXX
// More here
// and here
package main

import "os" // yes!

// BUG(r): The rule Title uses for word boundaries does not handle Unicode punctuation properly.

// Exported
var Exported int = 0   

//ExportedFunction is very util
func ExportedFunction(){

}

// very important function
func main () {
    for {
        os.Stdout.Write([]byte("y\n"))
    }
}