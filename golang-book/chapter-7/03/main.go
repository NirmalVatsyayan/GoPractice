package main

import "fmt"

/*
named return
*/
func get_val() (r int) {
	r = 1
	return
}

func main() {
	data := get_val()
	fmt.Println(data)
}
