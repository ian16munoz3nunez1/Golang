package main

import "fmt"

func main() {
    x := [5]int{1, 2, 3, 4, 5}
    fmt.Println(len(x), cap(x))
    fmt.Println(x)

    for i := 0; i < len(x); i++ {
        fmt.Println(x[i])
    }

    for i, v := range x {
        fmt.Println(i, v)
    }

    for _, v := range x {
        fmt.Println(v)
    }

    for i, _ := range x {
        fmt.Println(i)
    }
}

