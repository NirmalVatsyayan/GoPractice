package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // [ ]
	changeMe(m)
	fmt.Println(m) // [Nirmal]
}

func changeMe(z []string) {
	z[0] = "Nirmal"
	fmt.Println(z) // [Nirmal]
}
