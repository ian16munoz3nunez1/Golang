package main

import "fmt"

func main() {
    materias := map[string]int {
        "Programacion": 100,
        "Logica": 90,
        "Algoritmia": 95,
        "Calculo": 90,
    }

    fmt.Println(materias)
    fmt.Println(materias["Algoritmia"])
    materias["Algoritmia"] = 98
    fmt.Println(materias["Algoritmia"])

    v, err := materias["Calculo"]
    fmt.Println(v, err)
    v, err = materias["Contol"]
    fmt.Println(v, err)

    if v, err = materias["Logica"]; err {
        fmt.Println(v, err)
    }
}

