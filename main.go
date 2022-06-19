package main

import (
    "fmt"
    "os"
)

func main() {
    archivo, err := os.Create("go.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer archivo.Close()

    archivo.WriteString("Hola mundo!")
}

