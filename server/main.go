package main

import (
    "fmt"
    "net"
    "encoding/gob"
)

type Persona struct {
    Nombre string
    Email []string
}

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
    var persona Persona
    
    err := gob.NewDecoder(conexion).Decode(&persona)
    if err != nil {
        fmt.Println(err)
        return
    } else {
        fmt.Println(">", persona)
    }
}

func main() {
    go server()

    var input string
    fmt.Scanln(&input)
}

