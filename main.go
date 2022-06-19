package main

import (
    "encoding/json"
    "fmt"
)

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

    j, err := json.Marshal(materias)

    fmt.Println(err)
    fmt.Println(j)
    fmt.Println(string(j))
}

