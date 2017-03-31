package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("len - ", len(s), " cap - ", cap(s))

	s[0] = "n"
	s[1] = "v"
	s[2] = "d"
	fmt.Println("s -> ", s)
	fmt.Println("s[0] -> ", s[0])

	s = append(s, "x")
	fmt.Println(s)
	s = append(s, "y", "z")
	fmt.Println(s)

	// copying a slice to another slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	// slice of slice
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	// initialize a slice in single line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// 2d slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
