package main

import "fmt"

// package level scope
var name string = "Nirmal Vatsyayan"

func main() {
	fmt.Printf("Inside main , name is %v\n", name)
	foo()
}

func foo() {
	fmt.Printf("inside foo, name is %v\n", name)
}
