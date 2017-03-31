package main

import "fmt"

//traditional add function
func add(a int, b int) int {
	return a + b
}

//function returning multiple variables
func retMultVal() (int, int, int) {
	return 1, 2, 3
}

//variadic function
func variadicFunc(x ...int) int {
	var sum int
	for _, val := range x {
		sum = sum + val
	}
	return sum
}

func main() {
	// plain traditional add function
	sum := add(100, 100)
	fmt.Println(sum)

	// returning multiple values from function
	x, y, z := retMultVal()
	fmt.Println(x, " >>> ", y, " >>>> ", z)

	// variadic function
	vSum := variadicFunc(1, 2, 3)
	fmt.Println(vSum)
	vSum = variadicFunc(1, 2)
	fmt.Println(vSum)

	nums := []int{1, 2, 3, 4}
	vSum = variadicFunc(nums...) // passing slice to variadicFunc
	fmt.Println(vSum)
}
