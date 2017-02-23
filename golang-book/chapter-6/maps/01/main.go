package main

import "fmt"

/*
Unordered collection of key-value pairs
*/

func main() {
	fmt.Println("Working on maps")
	//creating maps

	x := make(map[string]string)
	x["name"] = "nirmal"
	x["surname"] = "vatsyayayn"

	fmt.Println(x)
	fmt.Println(x["name"])

	// length of map
	fmt.Println(len(x))

	//deleting a key from map
	delete(x, "name")
	fmt.Println(x)

	//looking for non existing key -> gives blank
	fmt.Println(x["name"])

	// getting value in one go and its status
	value, exists := x["surname"]
	fmt.Println(value, "  >>>>  ", exists)

	value1, exists1 := x["random"]
	fmt.Println(value1, "  >>>>  ", exists1)

	// graceful way

	if surname, existing := x["surname"]; existing {
		fmt.Println(surname)
	}

}
