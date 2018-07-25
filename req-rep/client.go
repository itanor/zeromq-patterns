package main

import (
  "fmt"
  zmq "github.com/alecthomas/gozmq"
)

func main() {
  context, _ := zmq.NewContext()
  socket, _ := context.NewSocket(zmq.REQ)
  defer context.Close()
  defer socket.Close()

  fmt.Printf("Connecting server...")
  socket.Connect("tcp://localhost:5555")

  for i := 0; i < 10; i++ {
    msg := fmt.Sprintf("Hello %d", i)
    socket.Send([]byte(msg), 0)
    println("Sending ", msg)

    reply, _ := socket.Recv(0)
    println("Received ", string(reply))
  }
}

