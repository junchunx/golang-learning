package main

import (
    "fmt"
    "example/newmath"
)

func main() {
    var x int
    fmt.Print("Input a integer: ")
    fmt.Scanln(&x)
    fmt.Printf("Sqrt(%d) = %v\n", x, newmath.Sqrt(float64(x)))
}

