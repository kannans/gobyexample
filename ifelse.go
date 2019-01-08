package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 6%2 != 0 {
		fmt.Println(" 6 is odd")
	} else if 6%2 == 0 {
		fmt.Println("6 is even")
	} else {
		fmt.Println("no number")
	}

	if num := 8; num < 10 {
		fmt.Println("Yes Its")
	}
}
