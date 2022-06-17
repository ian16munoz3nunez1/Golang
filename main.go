package main

import "fmt"

func main() {
    s := make([]int, 5, 16)

    fmt.Println(len(s), cap(s))
    s = append(s, 1, 2, 3, 4, 5)
    fmt.Println(len(s), cap(s))

    for _, v := range s {
        fmt.Println(v)
    }
}

