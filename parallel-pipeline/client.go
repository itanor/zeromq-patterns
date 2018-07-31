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

    fmt.Printf("Connecting to server...")
    socket.Connect("tcp://localhost:5555")

    for {
        fmt.Print("type something: ")
        var line string
        fmt.Scanln(&line)

        msg := fmt.Sprintf(line)

        socket.Send([]byte(msg), 0)
        println("Sending ", msg)

        _, _ = socket.Recv(0)
    }
}

