package main

import "fmt"

/*
using range function
*/
func main() {
	var sum float64
	var x [5]float64
	x[0] = 99.9
	x[1] = 99.1
	x[2] = 99.0
	x[3] = 99.2
	x[4] = 99.3

	for _, value := range x {
		sum += value
	}

	fmt.Println("Sum -> ", sum)
	fmt.Println("Average -> ", sum/float64(len(x)))

}
