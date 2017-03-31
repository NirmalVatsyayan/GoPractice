package main

import "fmt"

/*
    chan <- // means channel for sending data
    <-chan chan for receiving data only
*/

func ping(pings chan<- string , msg string){
     pings <- msg
}

func pong(pings <-chan string, pongs chan<- string){
     msg := <-pings
     pongs<-msg
}

func main(){
     pings := make(chan string, 1)
     pongs := make(chan string, 1)

     ping(pings, "Passed Message")
     pong(pings, pongs)
     fmt.Println(<-pongs)

}
