package main

import "fmt"

func main() {
	fmt.Printf(
		"90 degree F = %.2f degree C \n",
		func(fa float64) float64 {
			return (fa - 32.0) * (5.0 / 9.0)
		}(90),
	)
}