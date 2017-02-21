package main

import "fmt"

func main() {
	fmt.Println("Enter a number between 0-9 :")
	var input int

	fmt.Scanf("%d", &input)

	switch input {
	case 0:
		fmt.Println("Input is 0")
	case 1:
		fmt.Println("Input is 1")
	case 2:
		fmt.Println("Input is 2")
	case 3:
		fmt.Println("Input is 3")
	case 4:
		fmt.Println("Input is 4")
	case 5:
		fmt.Println("Input is 5")
	case 6:
		fmt.Println("Input is 6")
	case 7:
		fmt.Println("Input is 7")
	case 8:
		fmt.Println("Input is 8")
	case 9:
		fmt.Println("Input is 9")
	default:
		fmt.Println("Invalid input")

	}

}
