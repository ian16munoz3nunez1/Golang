package main

import "fmt"

func main() {
    var opcion int

    fmt.Print("Elige una opcion: ")
    fmt.Scan(&opcion)

    switch opcion {
        case 1:
            fmt.Println("Opcion 1")
        case 2:
            fmt.Println("Opcion 2")
        case 3:
            fmt.Println("Opcion 3")
        default:
            fmt.Println("La opcion", opcion, "no es valida")
    }
}

