package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	fmt.Print("Write ", i, " as --> ")

	switch i {
	case 1:
		fmt.Println("One !!")
	case 2:
		fmt.Println("Two !!")
	case 3:
		fmt.Println("Three !!")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its weekend !!")
	default:
		fmt.Println("Weekday !!")
	}

	t := time.Now()

	// switch without an condition , could be used as if-else
	switch {
	case t.Hour() > 12:
		fmt.Println("Good afternoon !!")
	case t.Hour() == 12:
		fmt.Println("Good noon !!")
	default:
		fmt.Println("Good morning !!")
	}

	// switch on functions
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
