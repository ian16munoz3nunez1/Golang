package main

import "fmt"

func main() {
    x := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
    s := x[0:4]
    fmt.Println(len(s), cap(s))
    for _, v := range s {
        fmt.Println(v)
    }
}

