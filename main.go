package main

import "fmt"

func main() {
    var temp int

    fmt.Print("Ingresa la temperatura: ")
    fmt.Scan(&temp)

    switch {
        case temp < 0:
            fmt.Println("Esta helado")
        case temp >= 0 && temp < 16:
            fmt.Println("Esta frio")
        case temp >= 16 && temp < 25:
            fmt.Println("La temperatura esta perfecta")
        case temp >= 25 && temp < 32:
            fmt.Println("Esta caliente")
        default:
            fmt.Println("Esta ardiendo")
    }
}

