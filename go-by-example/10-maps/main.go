package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["1"] = 1
	m["2"] = 2
	fmt.Println("m : ", m)
	fmt.Println("len(m) : ", len(m))

	x := m["1"]
	fmt.Println("x : ", x)
	delete(m, "1")
	fmt.Println(m)

	// first value will be blank identifier second is present or not
	a, b := m["1"]
	fmt.Println(a)
	fmt.Println(b)

	// declare and initialize in 1 statement
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
