package main

import "fmt"

func sum(nums ...int) {

	fmt.Println("Numbers:", nums)
	sum := 0

	for _, n := range nums {
	 sum = sum + n
	}

	fmt.Println("Total:", sum)
}


func minus(nums ...int) {
	fmt.Println("Minus values from 1000...", nums)
	total := 1000
	for _, i := range nums {
		total = total - i 	
	}
	fmt.Println("Total:", total)
}


func main() {
	sum(1,3,4,5,5,6,6,7,7,7,7,7,7)
	sum(45, 46)
	sum(40, -10)

	nums := []int{1,3,4,5,6,6,7} 
	sum(nums...)

	minus(45,3,5)
}
