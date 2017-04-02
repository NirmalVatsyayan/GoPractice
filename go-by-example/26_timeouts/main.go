package main

import (
  "fmt"
  "time"
)

func main(){

    // making processing timeout on operation with channel1
    c1 := make(chan string)
    go func(){
        time.Sleep(time.Second*2)
        c1 <- "message c1"
    }()

    select {
    case res := <- c1:
      fmt.Println(res)
    case <-time.After(time.Second*1):
      fmt.Println("Timeout 1 !!")
    }


    // making processing fine with operation on channel2
    c2 := make(chan string)
    go func(){
        time.Sleep(time.Second*1)
        c2 <- "message c2"
    }()

    select {
    case res := <- c2:
      fmt.Println(res)
    case <-time.After(time.Second*2):
      fmt.Println("Timeout 2 !!")
    }


}
