package main

import (
  "fmt"
  "math"
)


type rect struct {
    width, height float64
}

func (r rect) area() float64 {
    return r.width * r.height
}

func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

type circle struct {
    radius float64
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

/*
the types which needs interface wrapper must have
implemented the method signatures provided in interface
*/
type geometry interface {
    area() float64
    perim() float64
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main(){
  r := rect{width: 3, height: 4}
  c := circle{radius: 5}

  fmt.Println("calling rectangle methods directly")
  fmt.Println(r.area())
  fmt.Println(r.perim(), "\n\n")

  fmt.Println("calling circle methods directly")
  fmt.Println(c.area())
  fmt.Println(c.perim(), "\n\n")

  fmt.Println("calling interfaces")
  measure(r)
  measure(c)

}
