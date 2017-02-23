package main

import "fmt"

func incr(x *int) {
	*x += 1
}

func main() {
	// new gives a fresh new pointer variable
	intPtr := new(int)
	fmt.Println(*intPtr)

	incr(intPtr)
	fmt.Println(*intPtr)

	incr(intPtr)
	fmt.Println(*intPtr)

}
