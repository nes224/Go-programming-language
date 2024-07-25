package main

import "fmt"

type Car struct {
	name        string
	door, price int
}

type Ev struct {
	feature         string
	battery, ranges int
	Car
}

func (c Car) Show() {
	fmt.Printf("Car name : %s, Door : %d, Price : %d", c.name, c.door, c.price)
}

func (e Ev) Ranges() float64 {
	return float64(e.ranges) * 1.60934
}

func main() {
	m := Ev{
		feature: "FSD",
		battery: 250,
		ranges:  272,
		Car: Car{
			name:  "Tesla Model 3 RWD",
			door:  3,
			price: 1759000,
		},
	}

	fmt.Printf("\nName : %s", m.Car.name)
	fmt.Printf("\nName: %s", m.name)
	fmt.Println()
	m.Car.Show()
	fmt.Println()
	m.Show()
	fmt.Printf("\nRange : %f kms", m.Ranges())
}
