package main

import (
	"fmt"
	zmq "github.com/alecthomas/gozmq"
	"time"
)

func main() {
	context, _ := zmq.NewContext()
	defer context.Close()
	socket, _ := context.NewSocket(zmq.REQ)
	defer socket.Close()
	socket.Connect("tcp://127.0.0.1:5000")

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		msg := fmt.Sprintf("msg %d", i)
		socket.Send([]byte(msg), 0)
		println("Sending", msg)
		data, _ := socket.Recv(0)
		println("Got", string(data))
	}
}
