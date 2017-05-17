package main

import (
	"fmt"
	"os"
)

func main(){
	//panic("error is here")
	fmt.Println("hello")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
