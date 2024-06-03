package main

import "fmt"

func sum[V int | float64](x V, y V) V { return x + y }

func main() {
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10.25, 20.5))
}
