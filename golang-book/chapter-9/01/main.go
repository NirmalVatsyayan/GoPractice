package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {

	c := Circle{x: 5, y: 5, r: 5}
	fmt.Println(c.area())

}
