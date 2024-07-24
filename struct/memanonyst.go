package main

import (
	"fmt"
)

func main() {
	project := struct {
		id    int
		name  string
		price int
	}{
		id:    1256,
		name:  "Camera Vlog",
		price: 26000,
	}
	fmt.Println(project)
}
