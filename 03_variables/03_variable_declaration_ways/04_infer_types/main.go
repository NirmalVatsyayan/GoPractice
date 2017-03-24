package main

import "fmt"

func main() {

	var message = "Hello World!"

	// infer mixes type
	var a, b, c = 1, false, 3

	fmt.Println(message, a, b, c)
}
