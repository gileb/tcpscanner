package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main() {

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100

	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	a[2] = 5
	fmt.Println("put a[2]=5", a)

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	ab := make([]int, 5)
	printSlice("ab", ab)

	bc := make([]int, 0, 65535)
	bc = bc[:23]
	bc[22] = 33
	printSlice("bc", bc)

	cd := bc[:2]
	printSlice("cd", cd)

	de := cd[2:5]
	printSlice("de", de)
}
