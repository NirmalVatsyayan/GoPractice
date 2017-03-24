// hello_world1.go
package main

/*
Compile this first Go-program with: 6g test.go
This compiles to a file: test.6
which is linked with the command: 6l test.6
This produces the executable named: 6.out
which executes with the command: ./6.out
and produces the output: Hello, world

*/

func main() {
	println("Hello", "world")
}
