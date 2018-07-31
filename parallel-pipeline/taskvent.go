package main

import (
    "fmt"
    zmq "github.com/alecthomas/gozmq"
    "time"
)

func main() {
    context, _ := zmq.NewContext()
    defer context.Close()

    sender, _ := context.NewSocket(zmq.PUSH)
    defer sender.Close()
    sender.Bind("tcp://*:5557")

    fmt.Print("Press Enter when the workers are ready: ")

    var line string
    fmt.Scanln(&line)

    fmt.Println("Sending tasks to workers...")

    //  Socket to send start of batch message on
    rep, _ := context.NewSocket(zmq.REP)
    defer rep.Close()
    rep.Bind("tcp://*:5555")

    for {
      fmt.Println("esperando envio...")
      env, _ := rep.Recv(0);
      fmt.Println("recebeu " + string(env))

      fmt.Println(time.Now().UnixNano())

      msg := fmt.Sprintf(string(env))
      sender.Send([]byte(msg), 0)

      rep.Send([]byte(""), 0)
    }
}
