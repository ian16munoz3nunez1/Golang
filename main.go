package main

import "fmt"

func main() {
    materias := map[string]map[string]int {
        "Semestre 1": {
            "Programacion": 100,
            "Logica": 90,
        },
        "Semestre 2": {
            "Algoritmia": 95,
            "Calculo": 90,
        },
    }

    fmt.Println(materias)

    fmt.Println(materias["Semestre 1"])
    fmt.Println(materias["Semestre 2"]["Algoritmia"])
    materias["Semestre 2"]["Algoritmia"] = 98
    fmt.Println(materias["Semestre 2"]["Algoritmia"])
}

