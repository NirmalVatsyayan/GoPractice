package main

import "fmt"
import m "github.com/NirmalVatsyayan/GoPractice/golang-book/chapter-11/math"

/*
Above way of importing package is way of using ALIAS
we need to do m.Average.

If we dont use it, we will have to do math.Average

*/

func main() {
	fmt.Println("Testing packages in GO !!")
	val := []float64{1, 2, 3, 4}
	avg := m.Average(val)
	fmt.Println("Average -> ", avg)
}
