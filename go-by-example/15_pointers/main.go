package main

import "fmt"

func callByValue(val int) {
	val = 0
}

func callByReference(val *int) {
	*val = 0
}

func main() {
	var data = 100
	fmt.Println("Before call by value data ->> ", data)
	callByValue(data)
	fmt.Println("After call by value data ->> ", data)
	callByReference(&data)
	fmt.Println("After call by reference data ->> ", data)
}
