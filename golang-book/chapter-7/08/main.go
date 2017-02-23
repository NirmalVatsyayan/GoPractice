package main

import "fmt"

func first() {
	fmt.Println("Inside first")
}

func second() {
	fmt.Println("inside second")
}

/* defer, panic, recover */
func main() {
	// normal sequential flow
	fmt.Println("Normal flow")
	fmt.Println()
	first()
	second()

	// using defer
	fmt.Println("\nTesting usage of defer")
	fmt.Println()
	defer first()
	second()

	//using panic and recover
	defer func() {
		str := recover()
		fmt.Println(">>>>>>>>> ", str)
	}()

	panic("PANIC :p")

}
