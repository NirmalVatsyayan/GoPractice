package main

import (
	"fmt"

	"github.com/NirmalVatsyayan/GoPractice/thewaytogo/code_examples/chapter_10/person"
)

func main() {
	p := new(person.Person)
	// error: p.firstName undefined (cannot refer to unexported field or method firstName)
	// p.firstName = "Eric"
	p.SetFirstName("Eric")
	fmt.Println(p.FirstName()) // Output: Eric
}
