package main

import "fmt"

type product1 struct {
	Name string
	Cat  string
}

type product2 struct {
	Name string
	Cat  string
}

func (p product2) String() string {
	return fmt.Sprintf("%s <%s>", p.Name, p.Cat)
}

func main() {
	sk1 := &product1{
		Name: "Photo Retouch",
		Cat:  "Graphics",
	}
	fmt.Printf("p1: type: %T, value: %s\n\n", sk1, sk1)
	sk2 := product2{
		Name: "PHP Language",
		Cat:  "Programming",
	}
	s := sk2.String()
	fmt.Printf("p2: type: %T, value: %s\n\n", s, s)
}
