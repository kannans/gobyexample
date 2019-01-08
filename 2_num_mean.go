package main

import (
	"fmt"
)

func main() {
	// var x int
	// var y int
	x := 1
	y := 301

	fmt.Printf("X=%v, type of %T\n", x, x)
	fmt.Printf("X=%v, type of %T\n", y, x)

	var mean int
	mean = x + y/2
	fmt.Printf("result: %v, type of %T\n", mean, mean)
}
