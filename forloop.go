package main

import "fmt"

func main(){
	
	i :=1
	for i <= 3 {
		fmt.Println("Print number:", i)
		i = i + 1
	}

	for j := 7; j <=9; j++ {
		fmt.Println("J loop:", j)
	} 
	
	for {
		fmt.Println("loop")
		break;	
	}

	for k :=0; k <= 10; k++ {
		
		if k % 2 != 0 {
			continue;
		} 
		fmt.Println("Odd Number: ", k)
		
	}
}
