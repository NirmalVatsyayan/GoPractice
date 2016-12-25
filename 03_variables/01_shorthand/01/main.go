package main

import "fmt"

func main() {
	iteration := 10
	language := "golang"
	marks := 4.17
	is_ok := false

	fmt.Println(iteration)
	fmt.Println(language)
	fmt.Println(marks)
	fmt.Println(is_ok)

	fmt.Printf("iteration is %v and is of type %T\n", iteration, iteration)
	fmt.Printf("language is %v and is of type %T\n", language, language)
	fmt.Printf("marks is %v and is of type %T\n", marks, marks)
	fmt.Printf("is it enough learning ? %v and the answer is of type %T\n", is_ok, is_ok)

}
