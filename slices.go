
package main
import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("s: ", s)

	s[0] = "1"
	s[1] = "2"
	s[2] = "3"

	fmt.Println("s: ", s)
	fmt.Println("s len: ", len(s))

	s = append(s, "f")
	s = append(s, "e", "d")

	fmt.Println("s: ", s)

	c := make([]string, len(s))
	copy (c, s)

	fmt.Println("c: ", c)

	var a[7]int

	fmt.Println("a: ", a)

	b := [2] int { 1, 3 }
	fmt.Println("b: ", b)

	l := s[2:5]

	fmt.Println("ls: ", l)

	fmt.Println("ls: ", s[1:4])
	fmt.Println("ls: ", s[3:])

	fmt.Println("ls: ", s[5:])

	fmt.Println("ls: ", s[:4])

	fmt.Println("ls: ", s[:2])

	str := []string{"g", "d", "a"}

	fmt.Println("str: ", str)

	twoD := make([][]int, 13)
    for i := 0; i < 13; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}