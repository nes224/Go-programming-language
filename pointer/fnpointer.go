package main

import "fmt"

func main() {
	m := "NesTester"
	fmt.Println("Before : ", m)
	change(&m)

	fmt.Println("After : ", m)

}

func change(change *string)  {
	*change = "TesterNes"
}