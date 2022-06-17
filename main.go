package main

import "fmt"

func main() {
    x := [...]int{1, 2, 3, 4, 5}
    fmt.Println(len(x), cap(x))
    for _, v := range x {
        fmt.Println(v)
    }
}

