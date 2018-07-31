package main

import (
    "fmt"
    zmq "github.com/alecthomas/gozmq"
)

func main() {
    context, _ := zmq.NewContext()
    defer context.Close()

    receiver, _ := context.NewSocket(zmq.PULL)
    defer receiver.Close()
    receiver.Connect("tcp://localhost:5557")

    for {
        msgbytes, _ := receiver.Recv(0)
        fmt.Printf("%s\n", string(msgbytes))
    }
}
