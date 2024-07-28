package main

import "fmt"

func main() {
	sl := []int{8,6,-3,5,2,-8,10,-4,9}

	ch := make(chan int)

	go sum(sl[:3], ch)

	go sum(sl[3:6],ch)

	go sum(sl[6:],ch)

	x,y,z := <- ch, <-ch, <-ch
	fmt.Println(x,y,z,x+y+z)
}

func sum(sl []int, ch chan int) {
	sum := 0
	for _, v := range sl {
		sum += v 
	}

	ch <- sum
}