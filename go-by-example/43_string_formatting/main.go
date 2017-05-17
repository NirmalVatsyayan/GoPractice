package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main(){
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	// the %+v variant will include the struct’s field names
	fmt.Printf("%+v\n", p)

	// prints source code/snippet which produces the value
	fmt.Printf("%#v\n", p)

	// prints type
	fmt.Printf("%T\n", p)

	fmt.Printf("%t\n", true) // boolean

	fmt.Printf("%d\n", 123) // decimal representation

	fmt.Printf("%b\n", 14) // binary representation

	fmt.Printf("%c\n", 33) // character at

	fmt.Printf("%x\n", 456) // hex encoding

	fmt.Printf("%f\n", 78.9) // float

	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	fmt.Printf("%s\n", "\"string\"")

	fmt.Printf("%x\n", "hex this")

	fmt.Printf("%p\n", &p) // representation of pointer
	fmt.Printf("|%6d|%6d|\n", 12, 345) // width and precision of number to print
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}
