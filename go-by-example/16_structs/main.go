package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func main() {
    fmt.Println(Person{"nirmal", "vatsyayan"})
    fmt.Println(Person{FirstName:"nirmal"})
    fmt.Println(&Person{FirstName:"nirmal", LastName:"V"})

    sp := &Person{"nirmal", "vatsyayan"}
    fmt.Println(sp.FirstName)
    fmt.Println(sp.LastName)
}
