package main

import "fmt"

type mile float64
type kilometer float64

func (m mile) kilometers() kilometer {
	return kilometer(m * 1.60934)
}

func main() {
	var m mile = 5.0
	var k kilometer = 0.0
	k = m.kilometers()
	fmt.Println(m , "mile = ", k, "kilometers")
}