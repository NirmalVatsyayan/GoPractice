package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // now no routine etc can put any value in channel
	}()

	// range could be used only on closed channel
	for n := range c {
		fmt.Println(n)
	}
}
