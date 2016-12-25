package main

import "fmt"

func main() {
	s := greet("Jane ", "Doe")
	fmt.Println(s)
	fmt.Println(greet("Nirmal ", "Vatsyayan"))
}

func greet(fname, lname string) string {
	//sprint prints to string instead of stdout, it will return
	//a string with spaces between operands
	return fmt.Sprint(fname, lname)
}
