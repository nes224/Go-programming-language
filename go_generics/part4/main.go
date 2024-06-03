package main

import "fmt"

type GameOrMovie interface {
	getPrice() float64
}

type Game struct {
	Title    string
	Price    float64
	Platform string
}

type Movie struct {
	Title string
	Price float64
}

// Implment a method for each obj
func (obj Game) getPrice() float64 {
	return obj.Price
}

func (obj Movie) getPrice() float64 {
	return obj.Price
}

func Average[V GameOrMovie](obj []V) float64 {
	var result float64
	for i := range obj {
		result += obj[i].getPrice()
	}
	return result / float64(len(obj))
}

func main() {
	games := []Game{
		{
			Title:    "MINECRAFT",
			Price:    29.99,
			Platform: "PC",
		},
		{
			Title:    "FALLOUT 4",
			Price:    10.99,
			Platform: "PC",
		},
		{
			Title:    "NO MAN'S SKY",
			Price:    19.99,
			Platform: "PC",
		},
	}

	movies := []Movie{
		{
			Title: "Hello World",
			Price: 5.99,
		},
		{
			Title: "Death Not",
			Price: 6.99,
		},
		{
			Title: "Forrest Gump",
			Price: 7.99,
		},
	}

	fmt.Printf("Games: %.2f\n", Average(games))
	fmt.Printf("Movie: %.2f\n", Average(movies))

}
