package main

import "fmt"

func maps() {
	language := map[int]string{
		3: "English",
		4: "French",
		5: "Spanish",
	}

	var products = make(map[int]string)
	products[1] = "chair"
	products[2] = "table"

	for i, value := range language {
		fmt.Printf("language : %d %s\n",i, value)
	}
	fmt.Printf("Products: %v\n", products)
	fmt.Printf("Product with key 2 %s\n", products[2])
	delete(products,1)
	fmt.Println("Products", products)
}
