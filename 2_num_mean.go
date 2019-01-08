package main

import (
	"fmt"
)

func main() {
	x := 4.0
	y := 31.0

	fmt.Printf("X=%v, type of %T\n", x, x)
	fmt.Printf("X=%v, type of %T\n", y, x)

	mean := x + y/2.0
	fmt.Printf("result: %v, type of %T\n", mean, mean)
}
