package main

import (
    "fmt"
    "time"
)

func cont(numero int) {
    for i := 1; i <= 16; i++ {
        fmt.Println(numero, ":", i)
        time.Sleep(time.Millisecond * 200)
    }
}

func main() {
    go cont(0)
    go cont(1)

    var input string
    fmt.Scanln(&input)
}

