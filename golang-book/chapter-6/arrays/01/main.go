package main

import "fmt"

func main(){
    var x[5]float64
    var sum float64
    
    x[0] = 90.1
    x[1] = 91.9
    x[2] = 92.1
    x[3] = 93.2
    x[4] = 94.3
    //x[100]=100  -- error || invalid array inddex 100 || out of bounds
    fmt.Println(x)
    
    for i:=0; i<len(x); i++ {
        sum += x[i]
    }
    fmt.Println(sum)
    // type conversion
    fmt.Println("Avegrage -> ", sum/float64(len(x)))
    
}
