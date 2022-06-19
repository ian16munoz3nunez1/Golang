package main

import (
    "fmt"
    "os"
)

func main() {
    archivo, err := os.Open("go.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer archivo.Close()

    stat, err := archivo.Stat()
    if err != nil {
        fmt.Println(err)
        return
    }

    total := stat.Size()
    b := make([]byte, total)

    count, err := archivo.Read(b)
    if err != nil {
        fmt.Println(err)
        return
    }

    cadena := string(b)
    fmt.Println("bytes:", count)
    fmt.Println(cadena)
}

