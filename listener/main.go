package main

import (
  "fmt"
  "net"
)

func main(){
  conn, err := net.ListenPacket("udp","127.0.0.1:7777")
  if err != nil{
    fmt.Printf("[Error] %s\n",err)
  } else {
    defer fmt.Printf("Program Ended")
    defer conn.Close()
    fmt.Printf("[Info] Started Packet Listener\n")
    for {
      data := make([]byte,1024)
      n,addr,err := conn.ReadFrom(data)
      if err != nil{
        fmt.Printf("[Error] %s\n",err)
      } else {
        fmt.Printf("[%s] %s\n",addr,data[0:n])
      }
    }
  }
}

