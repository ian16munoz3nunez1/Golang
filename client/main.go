package main

import (
    "fmt"
    "net"
)

func client() {
    sock, err := net.Dial("tcp", ":9999")
    if err != nil {
        fmt.Println(err)
        return
    }

    msg := "Hola mundo!"
    fmt.Println(msg)
    sock.Write([]byte(msg))
    sock.Close()
}

func main() {
    go client()

    var input string
    fmt.Scanln(&input)
}

