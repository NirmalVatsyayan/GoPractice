package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for {
		c <- "ping"
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for {
		c <- "pong"
		time.Sleep(time.Second * 1)
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		time.Sleep(time.Second * 1)
		fmt.Println(msg)
	}
}

func main() {
	fmt.Println("Testing channels")

	channel := make(chan string)
	go pinger(channel)
	go printer(channel)
	go ponger(channel)
	var input string
	fmt.Scanln(&input)

}
