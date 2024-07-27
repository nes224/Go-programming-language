// in golang polymorphism == interface
package main

import (
	"fmt"
	"math"
)

// polymorphism คือ object หนึ่งสามารถตอบสนองการเรียกใช้งานได้หลากหลาย
type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length float64
	width float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rectangle) Area() float64 {
	return r.length + r.width
}

func totalArea(shapes []Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}

func main() {
	r1 := Rectangle{length: 10, width: 5}
	r2 := Rectangle{length: 4, width: 3}

	c1 := Circle{radius: 10}
	c2 := Circle{radius: 8}
	AllCircle := []Shape{c1,c2}

	Allrectangle := []Shape{r1,r2}
	Allshaps := []Shape{c1,c2,r1,r2}
	fmt.Printf("Total Area of cirlce = %.2f\n", totalArea(AllCircle))
	fmt.Printf("Total Area of rectangle = %.2f\n", totalArea(Allrectangle))
	fmt.Printf("Total Area of circle and rectangle = %.2f\n", totalArea(Allshaps))
}