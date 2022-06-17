package main

import "fmt"

func main() {
    var s []int

    fmt.Println(len(s), cap(s))
    s = append(s, 1, 2, 3, 4, 5)
    fmt.Println(len(s), cap(s))

    for _, v := range s {
        fmt.Println(v)
    }
}

