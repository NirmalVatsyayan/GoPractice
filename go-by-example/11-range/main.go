package main

import "fmt"

func main() {

	//declare initialize a slice
	nums := []int{2, 3, 4}
	sum := 0

	//ignore the index
	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// iterate over map key valye pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// iterate over keys of map
	for k := range kvs {
		fmt.Println("key:", k)
	}

	//iterate over runes in a string
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
