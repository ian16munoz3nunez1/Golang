package figuras

import "math"

type Figura interface {
    Area() float64
}

type Circulo struct {
    Radio float64
}

func (circulo *Circulo) Area() float64 {
    return circulo.Radio * circulo.Radio * math.Pi
}

type Rectangulo struct {
    Altura float64
    Base float64
}

func (rectangulo *Rectangulo) Area() float64 {
    return rectangulo.Altura * rectangulo.Base
}

