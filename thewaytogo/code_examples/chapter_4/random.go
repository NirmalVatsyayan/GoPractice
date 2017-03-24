package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d / ", a)
	}
	for i := 0; i < 5; i++ {
		r := rand.Intn(8)
		fmt.Printf("%d / ", r)
	}
	fmt.Println()
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32())
	}
}

/* Output:
5577006791947779410 / 8674665223082153551 / 6129484611666145821 / 4037200794235010051 / 3916589616287113937 / 6334824724549167320
/ 605394647632969758 / 1443635317331776148 / 894385949183117216 / 2775422040480279449 / 6 / 7 / 2 / 1 / 0 /
97.88 / 21.68 / 88.43 / 85.96 / 68.27 / 46.43 / 18.89 / 59.49 / 81.47 / 49.07 /
*/
