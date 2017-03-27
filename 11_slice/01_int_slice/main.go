package main

import "fmt"

func main() {

	//https://gobyexample.com/slices
	fmt.Println(len([]int{}))
	fmt.Println(cap([]int{}))

	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Println(cap(mySlice))
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
	mySlice = append(mySlice, 12)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))

	fmt.Println(mySlice[2])   // index access
	fmt.Println(mySlice[2:4]) // slicing a slice

}
