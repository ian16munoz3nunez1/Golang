package main

import (
    "fmt"
    "container/list"
)

func main() {
    var lista list.List

    lista.PushBack(4)
    lista.PushBack("Go")
    lista.PushFront(16.84)
    lista.PushFront(true)

    for e := lista.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }

    fmt.Println(lista)
}

