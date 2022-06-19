package main

import (
    "fmt"
    "sort"
)

func main() {
    s1 := []int{8, 7, 6, 5}

    sort.Ints(s1)
    fmt.Println(s1)

    s2 := []int{1, 2, 3, 4}
    sort.Sort(sort.Reverse(sort.IntSlice(s2)))
    fmt.Println(s2)
}

