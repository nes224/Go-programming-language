package main

import "fmt"

type Game struct {
	Title    string
	Price    float64
	Platform string
}

type Movie struct {
	Title string
	Price float64
}

func AveragGamePrice(obj []Game) float64 {
	var result float64
	for i := range obj {
		result += obj[i].Price
	}
	return result / float64(len(obj))
}

func AverageMoviePrice(obj []Movie) float64 {
	var result float64
	for i := range obj {
		result += obj[i].Price
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
			Title: "Death Note",
			Price: 6.99,
		},
		{
			Title: "Forrest Gump",
			Price: 7.99,
		},
	}

	fmt.Printf("games: %.2f\n", AveragGamePrice(games))
	fmt.Printf("mvoies: %.2f\n", AverageMoviePrice(movies))

}
