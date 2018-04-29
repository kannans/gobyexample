package main

import "fmt"

func main() {
	var a[5]int
	a[2]=100
	fmt.Println(a[2])

	b :=[4]int{1,3,5}
	fmt.Println(b)

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
