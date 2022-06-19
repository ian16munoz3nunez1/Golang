package main

import (
    "fmt"
    "sort"
)

type Persona struct {
    nombre string
    edad uint64
}

type ByNombre []Persona
func (a ByNombre) Len() int {return len(a)}
func (a ByNombre) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
func (a ByNombre) Less(i, j int) bool {return a[i].nombre < a[j].nombre}

type ByEdad []Persona
func (a ByEdad) Len() int {return len(a)}
func (a ByEdad) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
func (a ByEdad) Less(i, j int) bool {return a[i].edad < a[j].edad}

func main() {
    ps := []Persona{
        Persona{"persona3", 31},
        Persona{"persona1", 20},
        Persona{"persona2", 54},
    }
    fmt.Println(ps)

    sort.Sort(ByNombre(ps))
    fmt.Println(ps)

    sort.Sort(ByEdad(ps))
    fmt.Println(ps)
}

