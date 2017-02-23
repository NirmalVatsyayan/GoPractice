package main

import "fmt"

/*
function to take variable length argument
*/
func variadic_example(args ...int) int {
	sum := 0
	for _, value := range args {
		sum += value
	}
	return sum
}

func main() {

	x := variadic_example(1, 2, 3, 4)
	fmt.Println(x)
}
