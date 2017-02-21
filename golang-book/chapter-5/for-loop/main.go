package main

import "fmt"

func main() {
	var x string

	x = "Syntax 1"
	i := 0
	for i < 5 {
		fmt.Println(x, i)
		i += 1
	}

	x = "Syntax 2"
	for j := 0; j < 5; j++ {
		fmt.Println(x, j)
	}

}
