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

func client(persona Persona) {
    sock, err := net.Dial("tcp", ":9999")
    if err != nil {
        fmt.Println(err)
        return
    }

    err = gob.NewEncoder(sock).Encode(persona)
    if err != nil {
        fmt.Println(err)
    }
    sock.Close()
}

func main() {
    persona := Persona{
        "juan",
        []string{
            "correo1@hotmail.com",
            "correo2@gmail.com",
        },
    }
    go client(persona)

    var input string
    fmt.Scanln(&input)
}

