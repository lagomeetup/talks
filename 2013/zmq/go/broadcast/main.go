package main

import (
	zmq "github.com/alecthomas/gozmq"
	"os"
	"strings"
)

func main() {
	context, _ := zmq.NewContext()
	defer context.Close()
	socket, _ := context.NewSocket(zmq.PUSH)
	defer socket.Close()

	socket.Connect("tcp://127.0.0.1:4000")

	msg := strings.Join(os.Args[1:], " ")
	socket.Send([]byte(msg), 0)
	println("Sending", msg)
}
