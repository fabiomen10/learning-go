package main

import (
        "fmt"
)

func main() {

        x := 16
        y := "strings"
        z := true
        
        _, erros := fmt.Println("Hello world!", "Oi galera!", 100)
        fmt.Println(erros)
        fmt.Println(x, y, z)
}
