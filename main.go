package main

import (
    "fmt"
    "figuras/figuras"
)

func main() {
    c01 := figuras.Circulo{16}
    r01 := figuras.Rectangulo{2, 8}

    fmt.Println(c01, c01.Area())
    fmt.Println(r01, r01.Area())
}

