package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // var numero float64
    scanner := bufio.NewScanner(os.Stdin)
    var nombre string

    // fmt.Print("Numero: ")
    // fmt.Scanf("%f", &numero)
    // fmt.Scan(&numero)
    fmt.Print("Ingresa tu nombre: ")
    scanner.Scan()
    nombre = scanner.Text()

    //multiplicacion := numero * 4
    //fmt.Println("Multiplicacion:", multiplicacion)
    fmt.Println("Hola", nombre)
}

