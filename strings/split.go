package main

import (
	"fmt"
	"strings"
)

func main() {
	ar1 := strings.Split("tester1 tester2, tester777@gmail.com, BKK", ",")
	ar2 := strings.Split("tester3 tester4 tester888@gmail.com BKK2", " ")
	fmt.Printf("%q\n", ar1)
	fmt.Printf("%q\n", ar2)
}