package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}


func multiRtr() (int, int) {
	return 4, 5
}


func main() {
	sum := plus(2,4)
	fmt.Println(sum)

	mins := minus(38,12)
	fmt.Println(mins)

	a, b := multiRtr()
	fmt.Println(a)
	fmt.Println(b)
}
