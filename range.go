package main
import "fmt"

func main() {
	nums := []int{1,2,3}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println("Val", sum)
	
	for i, num := range nums {
		fmt.Println("Index:", i, "Num:", num)
	}

	names := map[string]string{"1": "Kannan", "2": "Suhas", "3": "github"}

	for key, value := range names {
		fmt.Println("My name is:", value, "ID:", key)
	}

	for k, v := range names {
		fmt.Printf("%s -> %s \n", k, v)
	}
	
}
