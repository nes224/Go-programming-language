// Deadlock

package main 
import "fmt"

func main() {
	ch := make(chan string)
	select {
	case msg := <-ch:
		fmt.Println(msg)
	default:
		fmt.Println("not receiving message from goroutines.")
	}
}