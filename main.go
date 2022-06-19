package main

import "fmt"

func main() {
    materias := make(map[string]int)
    materias["Programacion"] = 100
    materias["Logica"] = 90
    materias["Algoritmia"] = 95
    materias["Calculo"] = 90

    fmt.Println(materias)
}

