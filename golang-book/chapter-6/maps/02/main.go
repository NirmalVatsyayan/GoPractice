package main

import "fmt"

//creating maps of maps
func main() {
	fmt.Println("creating maps of maps")

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "hydrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "oxygen",
			"state": "gas",
		},
	}

	if element, ok := elements["O"]; ok {
		fmt.Println(element)
	}

}
