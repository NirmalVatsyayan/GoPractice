package main

import "fmt"

/*
create a slice and append / copy value in it
*/
func main() {
	// method 1
	x := []int{1, 2, 3}
	y := append(x, 4, 5, 6)

	fmt.Println(x)
	fmt.Println(y)

	// method 2
	z := x[0:3]
	fmt.Println(z)
	a := append(z, 4, 5, 6)
	fmt.Println(a)

	//method 3
	b := make([]int, 5)
	c := append(b, 1, 2, 3)
	fmt.Println(c)

	// method 4
	slice1 := []int{1, 2, 3, 4}
	fmt.Println(slice1)
	slice2 := make([]int, 4)
	copy(slice2, slice1)
	fmt.Println(slice2)

}
