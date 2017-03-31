package main

import "fmt"

/*
VVI
Go automatically handles conversion between values and pointers for method calls.
You may want to use a pointer receiver type to avoid copying on method calls or
to allow the method to mutate the receiving struct.
*/

type Rect struct{
  Width int
  Height int
}

/*
Methods can be defined for pointer or value receiver types
*/
func (r Rect) area() int {
  return r.Width * r.Height
}

func (r *Rect) perim() int{
  return 2*r.Width + 2*r.Height
}

func main(){
  r := Rect{Width:10, Height:3}
  fmt.Println(r.area())
  fmt.Println(r.perim()) // no need to pass address , its handled by go itself

  sp := &r
  fmt.Println(sp.area())
  fmt.Println(sp.perim())

}
