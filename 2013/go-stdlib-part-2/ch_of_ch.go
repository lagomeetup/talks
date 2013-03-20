package main


type ch2 chan chan interface{}


func main() {
	mych2 := make(ch2,1)
	ch := make(chan interface{},1)
	ch <- 1
	mych2 <- ch
	<- <- mych2
	close(mych2)
}
