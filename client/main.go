package main

import (
    "fmt"
    "net"
    "encoding/gob"
)

func client() {
    sock, err := net.Dial("tcp", ":9999")
    if err != nil {
        fmt.Println(err)
        return
    }

    msg := "Hola mundo!"
    fmt.Println(msg)
    err = gob.NewEncoder(sock).Encode(msg)
    if err != nil {
        fmt.Println(err)
    }
    sock.Close()
}

func main() {
    go client()

    var input string
    fmt.Scanln(&input)
}

