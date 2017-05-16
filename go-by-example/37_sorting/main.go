package main

import (
	"sort"
	"fmt"
)

func main(){
	strs := []string{"c", "a", "b"}
	fmt.Println(sort.StringsAreSorted(strs))
	sort.Strings(strs)
	fmt.Println(strs)
	fmt.Println(sort.StringsAreSorted(strs))

	ints := []int{5,3,1}
	fmt.Println(sort.IntsAreSorted(ints))
	sort.Ints(ints)
	fmt.Println(ints)
	fmt.Println(sort.IntsAreSorted(ints))

}
