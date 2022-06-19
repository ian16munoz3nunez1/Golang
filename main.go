package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func () {
        for {
            c1 <- "Mensaje desde el canal 1"
            time.Sleep(time.Second * 2)
        }
    }()

    go func () {
        for {
            c2 <- "Mensaje desde el canal 2"
            time.Sleep(time.Second)
        }
    }()

    go func () {
        for {
            select {
                case msg := <- c1:
                    fmt.Println(msg)
                case msg := <- c2:
                    fmt.Println(msg)
            }
        }
    }()

    var input string
    fmt.Scanln(&input)
}

