package main

import "fmt"

func average(x []float64) float64 {
	sum := 0.0

	for _, value := range x {
		sum += value
	}

	return sum / float64(len(x))

}

func main() {
	fmt.Println("Inside main function")
	x := average([]float64{1, 2, 3, 4})
	fmt.Println("average is -> ", x)
}
