package main

import (
    "fmt"
    "strings"
)

func main() {
    cadena := "Arduino-C-Cpp-CSS-Go-HTML-Java-JS-Latino-Python-Ruby"

    fmt.Println(cadena)
    fmt.Println(strings.Contains(cadena, "Go"))
    fmt.Println(strings.Count(cadena, "C"))
    fmt.Println(strings.HasPrefix(cadena, "Arduino"))
    fmt.Println(strings.HasSuffix(cadena, "Ruby"))
    fmt.Println(strings.Index(cadena, "Go"))
    fmt.Println(strings.Join([]string{"Android", "MATLAB", "MPLAB"}, "-"))
    fmt.Println(strings.Repeat("Go ", 4))
    fmt.Println(strings.Replace(cadena, "Python", "PHP", 1))
    fmt.Println(strings.Split(cadena, "-"))
    fmt.Println(strings.ToLower(cadena))
    fmt.Println(strings.ToUpper(cadena))

    b := []byte("go")
    fmt.Println(b)

    s := string([]byte{'g', 'o'})
    fmt.Println(s)
}

