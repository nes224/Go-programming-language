package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("small small small","ll", "ller", 2))
	fmt.Println(strings.Replace("small small small","all", "ile", -1))
}
