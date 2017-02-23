package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func f(n int, name string) {
	for i := 0; i < n; i++ {
		fmt.Println("inside goroutine ", name, " i> ", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i, "routine "+strconv.Itoa(i))
	}

	var input string
	fmt.Scanln(&input)
}
