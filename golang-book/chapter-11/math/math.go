package math

import "fmt"

func Average(val []float64) float64 {
	fmt.Println("Inside Average function of MATH package")
	sum := 0.0
	for _, value := range val {
		sum += value
	}
	return sum / float64(len(val))
}
