package main

import "fmt"

func fact(val int) int {
	if val == 0 {
		return 1
	}

	return val * fact(val-1)
}

func main() {
	fact4 := fact(4)
	fmt.Println(fact4)
}
