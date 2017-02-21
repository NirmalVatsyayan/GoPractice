package main

import "fmt"

func main() {
	const x string = "Nirmal vatsyayan!!"
	fmt.Println(x)

	//x = "Nirmal" will give an error

	//defining multiple variables
	const (
		a = 10
		b = 100
		c = 1000
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
