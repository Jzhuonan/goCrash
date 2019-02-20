package main

import (
	"fmt"
	"math"
)

// Shaper interface
type Shaper interface {
	area() float64
}

// Circle struct
type Circle struct {
	x, y, radius float64
}

// Rectangle struct
type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height 
}

func getArea(s Shaper) float64 {
	return s.area()
}

func main() {

	c := Circle{0, 0, 5}
	r := Rectangle{width: 6, height: 4}

	fmt.Println(getArea(c))
	fmt.Println(getArea(r))
}