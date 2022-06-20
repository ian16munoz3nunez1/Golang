package main

import (
    "fmt"
    "net"
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
    msg := make([]byte, 100)

    b, err := conexion.Read(msg)
    if err != nil {
        fmt.Println(err)
        return
    } else {
        fmt.Println("Bytes:", b)
        fmt.Println(">", string(msg[:b]))
    }
}

func main() {
    go server()

    var input string
    fmt.Scanln(&input)
}

