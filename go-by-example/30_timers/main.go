package main

import (
	"time"
	"fmt"
)

func main()  {

	q := make(chan string, 2)

	timer1 := time.NewTimer(time.Second * 2)
	<- timer1.C
	fmt.Println("timer 1 expired")

	timer2 := time.NewTimer(time.Second)

	go func(){
		<- timer2.C
		q <- "timer 2 expired"
		close(q)
	}()



	stop := timer2.Stop()
	fmt.Print(stop)
	if stop {
		fmt.Println("timer 2 stopped")
	}

	if len(q)>0{
		fmt.Println(<-q)
	}

}
