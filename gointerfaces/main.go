package main

import "fmt"

type geometry interface {
	area() float64
	perim() float64
}
type rect struct {
	height, width float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func main() {
	r := rect{width: 3, height: 4}
	fmt.Println("area: ", r.area())
}
