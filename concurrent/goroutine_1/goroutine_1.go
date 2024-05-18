package main

import (
	"fmt"
	"strings"
	"time"
)

func processText(text string) {
	parts := strings.Split(text, "|")
	name := parts[0]
	age := parts[1]
	gender := parts[2]

	time.Sleep(3 * time.Second) 
	fmt.Printf("Send email >> Name=%-5s gender=%-6s age=%2s\n", name, gender, age)
}

func main() {
	allText := []string{
		"Perth|29|male",
		"Noom|25|male",
		"Onny|25|female",
	}

	for _, text := range allText {
		go processText(text)
		fmt.Println("Your request has been added to queue.")
	}

	time.Sleep(4 * time.Second)

}
