package main

import "fmt"

/*
array initialization and instantiation in 1 go
*/
func main() {
	x := [10]float64{90.0, 91.1, 92.2, 93.3, 94.4, 95.5, 96.6, 97.7, 98.8, 99.9}

	var sum float64

	for _, value := range x {
		sum += value
	}

	fmt.Println("Sum is -> ", sum)
	fmt.Println("Average is -> ", sum/float64(len(x)))

}
