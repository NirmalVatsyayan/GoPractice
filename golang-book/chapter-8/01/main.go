package main

import "fmt"

func change(x int) {
	x += 5
}

func changePtr(x *int) {
	// dereferencing
	*x = 5
}

// pointer
func main() {
	x := 0

	// default function calls are by value
	change(x)
	fmt.Println("x is now ", x)

	changePtr(&x)
	fmt.Println("x is now ", x)
}
