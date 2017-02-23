package main

import "fmt"

func return_multiple() (int, int, int) {
	return 1, 2, 3
}

func main() {
	a, b, c := return_multiple()
	fmt.Println(a, " >> ", b, " >> ", c)
}
