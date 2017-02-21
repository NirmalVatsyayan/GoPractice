package main

import "fmt"

// Booleans are named after George Boole
// 1 bit -> true or false i.e on or off
/*

&&  and
|| or
!  not

*/
func main() {
	fmt.Println(true && false)
	//fmt.Println(true && true)

	fmt.Println(true || false)
	//fmt.Println(true || true)

	fmt.Println(!false)
	fmt.Println(!true)

}
