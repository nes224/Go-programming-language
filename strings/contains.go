package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println(strings.Compare("a","z"))
	fmt.Println(strings.Compare("z","z"))
	fmt.Println(strings.Compare("z","a"))
}