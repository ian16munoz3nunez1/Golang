package main

import (
    "fmt"
    "figuras/figuras"
)

func sumAreas(figs ...figuras.Figura) float64{
    areaTotal := 0.0

    for _, fig := range figs {
        areaTotal += fig.Area()
    }

    return areaTotal
}

func main() {
    c01 := figuras.Circulo{16}
    r01 := figuras.Rectangulo{2, 8}
    fmt.Println(c01, c01.Area())
    fmt.Println(r01, r01.Area())

    fmt.Println(sumAreas(&c01, &r01))

    mf := figuras.MultiFigura{
        Figuras: []figuras.Figura{
            &c01,
            &r01,
            &figuras.Circulo{4},
        },
    }
    fmt.Println(mf.Area())
}

