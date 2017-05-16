package main

import "fmt"

func main(){
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) // we could close a non-empty channel as well

	for elem := range queue {
		fmt.Println(elem)
	}
}
