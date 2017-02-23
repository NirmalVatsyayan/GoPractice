package main

import "fmt"

func main(){
    var x[5]int
    var sum int
    
    x[0] = 90
    x[1] = 91
    x[2] = 92
    x[3] = 93
    x[4] = 94
    //x[100]=100  -- error || invalid array inddex 100 || out of bounds
    fmt.Println(x)
    
    for i:=0; i<len(x); i++ {
        sum += x[i]
    }
    fmt.Println(sum)
    fmt.Println("Avegrage -> ", sum/len(x))
    
}
