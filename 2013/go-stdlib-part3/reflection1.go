package main


func main() {
    var x interface{}	
    type MyInt int

    var i int 
    var j MyInt

    i = j // ERROR

}