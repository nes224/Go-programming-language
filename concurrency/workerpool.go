package main

import (
	"fmt"
	"time"
)

// We’re controlling concurrency while efficiently processing jobs. 
// It’s a lifesaver when dealing with things like handling multiple incoming requests or processing data batches.

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second)
		results <- j * 2
		fmt.Printf("Worker %d finished job %d\n", id, j)
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs<-j
	}
	close(jobs)
	for a := 1;a <= numJobs; a++ {
		<-results
	}
}
