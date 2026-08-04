package main

import (
  "fmt"
  "net"
)   
/*
const (
  MulticastGroup = "239.255.255.250"
  MulticastPort = "3702"
)*/

func main(){
  conn, err := net.Dial("udp", "127.0.0.1:7777")
  if err != nil{
    fmt.Printf("[Error]:%s\n",err)
  }else{
    defer conn.Close()
    message := []byte("Hello World")
    _, err := conn.Write(message)
    if err != nil{
      fmt.Printf("[Error]:%s\n",err)
    }else{
      fmt.Printf("Successfully ran the code\n")
    }
  }
}

