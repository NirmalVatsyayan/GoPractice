package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even !!")
	} else {
		fmt.Println("7 is odd !!")
	}

	if 8%2 == 0 {
		fmt.Println("8 is even !!")
	}

	if num := 9; num < 0 {
		fmt.Println(num, " is negative !!")
	} else if num < 10 {
		fmt.Println(num, " is of 1 digit !!")
	} else {
		fmt.Println(num, " has multiple digits !!")
	}

}
