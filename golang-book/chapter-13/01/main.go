package main

import (
	"fmt"
	"strings"
)

/*
using string library of golang
*/

func main() {
	fmt.Println("Inside main function")

	//checks occurance of substr in str
	fmt.Println(strings.Contains("vatsyayan", "yan"))

	//check count of occurance of sub str in str
	fmt.Println(strings.Count("vatsyayan", "y"))

	//check prefix and/or suffix
	fmt.Println(strings.HasPrefix("nirmal", "ni"))
	fmt.Println(strings.HasSuffix("vatsyayan", "an"))

	//index of occurance of a sub str
	fmt.Println(strings.Index("nirmal", "i"))
	fmt.Println(strings.Index("nirmal", "z"))

	//join strings
	fmt.Println(strings.Join([]string{"nirmal", "vatsyayan"}, "-"))

}
