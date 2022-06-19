package main

import (
    "fmt"
    "math"
)

type Circulo struct {
    radio float64
}

func (circulo *Circulo) area() float64 {
    return circulo.radio * circulo.radio * math.Pi
}

func main() {
    c01 := Circulo{}
    c02 := Circulo{radio: 16}
    c03 := Circulo{8}
    c04 := new(Circulo)
    c05 := &Circulo{4}

    fmt.Println(c01, c01.area())
    fmt.Println(c02, c02.area())
    fmt.Println(c03, c03.area())
    fmt.Println(c04, c04.area())
    fmt.Println(c05, c05.area())
}

