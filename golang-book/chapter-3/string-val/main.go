package main

import "fmt"

func main() {
	x := "Double quote: This is a string \n value"
	fmt.Println(x)

	y := `Back ticks: This is also a string \n value`
	fmt.Println(y)

	/*
	   Difference  between both are
	   Double quoted cannot contain new line and tab char,
	   they will be decomposed to take action but they also
	   allow special escape sequences

	   But in back ticks, they will be treated as simple char
	*/

	// finding length of string
	fmt.Println("Length of x is ", len(x), " length of y is ", len(y))

	// slicing string, it will give numeric value
	// as characters are represented as bytes
	fmt.Println("Data at first location of : x -> ", x[0], " and y -> ", y[0])

	// concatenation of strings
	fmt.Println("Hello " + " World !! " + " AGAIN :P ")

}
