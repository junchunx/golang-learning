package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	var sum bool
	flag.BoolVar(&sum, "sum", false, "sum of all values")
	flag.Parse()
	flag.Usage()
	fmt.Println("sum =", sum)
	fmt.Println("num of flags", flag.NFlag())
	fmt.Println(flag.Args())
	fmt.Println("num of args:", flag.NArg())
	fmt.Println("arg 1:", flag.Arg(1))
	total := 0
	for _, str := range flag.Args() {
		val, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Ignore invald value %v, must be integer\n", str)
		}
		total += val
	}
	fmt.Println(total)
}
