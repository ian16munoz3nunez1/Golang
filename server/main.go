package main

import (
    "fmt"
    "net"
    "encoding/gob"
)

func server() {
    sock, err := net.Listen("tcp", ":9999")
    if err != nil {
        fmt.Println(err)
        return
    }

    for {
        conexion, err := sock.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }

        go handleClient(conexion)
    }
}

func handleClient(conexion net.Conn) {
    var msg string
    
    err := gob.NewDecoder(conexion).Decode(&msg)
    if err != nil {
        fmt.Println(err)
        return
    } else {
        fmt.Println(">", msg)
    }
}

func main() {
    go server()

    var input string
    fmt.Scanln(&input)
}

