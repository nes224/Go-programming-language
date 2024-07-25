package main
import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	length int
	width int
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (r Rectangle) perimeter() int {
	return 2 * (r.length + r .width)
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	rec := Rectangle {
		length: 10,
		width: 5,
	}

	cir := Circle {radius: 10}

	fmt.Printf("Area of rectangle %d\n", rec.Area())
	fmt.Printf("Perimeter of rectangle %d\n", rec.perimeter())

	fmt.Printf("Area of circle %.2f\n", cir.Area())
	fmt.Printf("Perimeter of circle %.2f", cir.perimeter())
}