package main

import "fmt"

func showData(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println(name, " --> ", i)
	}
}

func main() {

	showData("sequential")
	go showData("goRoutine")

	//goRoutine for anonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("Nirmal")

	var input string // can use waitgroup instead of this
	fmt.Scanln(&input)
	fmt.Println("Done !!")
}
