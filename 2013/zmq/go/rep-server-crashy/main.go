package main

import zmq "github.com/alecthomas/gozmq"

func main() {
	context, _ := zmq.NewContext()
	defer context.Close()
	socket, _ := context.NewSocket(zmq.REP)
	defer socket.Close()
	socket.Bind("tcp://127.0.0.1:5000")

	for {
		msg, _ := socket.Recv(0)
		println("Got", string(msg))
		if string(msg) == "10" {
			panic("random crash")
		}
		socket.Send(msg, 0)
	}
}
