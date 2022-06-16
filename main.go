package main

import "fmt"

func main() {
    var temp int

    fmt.Print("Ingresa la temperatura: ")
    fmt.Scan(&temp)

    if temp < 0 {
        fmt.Println("Esta helado")
    } else if temp >= 0 && temp < 16 {
        fmt.Println("Esta frio")
    } else if temp >= 16 && temp < 25 {
        fmt.Println("La temperatura esta perfecta")
    } else if temp >= 25 && temp < 32 {
        fmt.Println("Esta caliente")
    } else {
        fmt.Println("Esta ardiendo")
    }
}

