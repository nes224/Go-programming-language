package main

import "fmt"

type product struct {
	int
	string
	float64
}

func main() {
	pd := product{1256, "CameraVlog", 0.45}
	fmt.Println("Detail of product")
	fmt.Println("Product ID : ", pd.int)
	fmt.Println("Product NAME : ", pd.string)
	fmt.Println("Product WEIGHT : ", pd.float64)
}