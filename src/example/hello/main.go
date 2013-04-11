package main

import (
	"example/newmath"
	"fmt"
)

func main() {
    var x int
	fmt.Print("Input a integer: ")
	fmt.Scanln(&x)
	fmt.Printf("Sqrt(%d) = %v\n", x, newmath.Sqrt(float64(x)))
}
