package main

import "fmt"

func main() {
    for i := 1; i <= 16; i++ {
        fmt.Println(i)
    }

    i := 16
    for i >= 1 {
        fmt.Println(i)
        i--
    }
}

