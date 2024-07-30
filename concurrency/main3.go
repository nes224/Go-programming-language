package main

import "fmt"

func InitServer(text string, ch chan string) {
	ch <- text
}

func main() {
	ch := make(chan string)
	buildText := []string{"I Love Golang", "Software Developer"}
	for i := range buildText {
		go InitServer(buildText[i],ch)
		select {
		case msg := <- ch:
			fmt.Println(msg)
			return
		default:
			fmt.Println("Nothing happen.")
		}
	}
}
