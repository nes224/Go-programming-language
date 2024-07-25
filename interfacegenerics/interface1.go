package main

import (
	"math"
	"fmt"
)

type Shape interface {
	Area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, length float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rectangle) Area() float64 {
	return float64(r.width) * float64(r.length)
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main() {
	var s Shape
	s = Circle{12}
	fmt.Printf("Area of circle : %.2f\n", s.Area())
	fmt.Printf("Perimeter of circle : %.2f\n", s.perimeter())

	s = Rectangle{6,4}
	fmt.Printf("Area of rectangle : %.2f\n", s.Area())
	fmt.Printf("Perimeter of area : %.2f\n", s.perimeter())
}
