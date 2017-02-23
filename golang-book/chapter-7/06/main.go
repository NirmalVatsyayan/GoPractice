package main

import "fmt"

/*
closure example
*/

func evenNumGenerator() func() uint {
	x := uint(0)
	return func() (ret uint) {
		ret = x
		x += 2
		return
	}

}

func main() {
	fmt.Println("creating closure")

	x := 0

	add := func(x int, y int) int {
		return x + y
	}

	fmt.Println(add(1, 2))

	// scope of local variable test
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())

	// even number generator test
	evenNum := evenNumGenerator()
	fmt.Println(evenNum())
	fmt.Println(evenNum())

}
