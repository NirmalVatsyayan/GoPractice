package main

import "fmt"

//variable defined in outer score of inner function
//is available till life time of inner function
func closureFunc() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextVal := closureFunc()
	fmt.Println(nextVal())
	fmt.Println(nextVal())
	fmt.Println(nextVal())
}
