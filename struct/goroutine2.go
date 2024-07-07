package main

import (
	"fmt"
	"net/http"
	"sync"
)

// every single thread consumes ~1MB of memory. Thus if your application creates 1000 threads then it would consume ~1GB, which is huge.
// light weight process. light weight threads.
// Goroutines are managed by Go Runtime. Every single goroutine consume ~2KB, https://miro.medium.com/v2/resize:fit:720/format:webp/1*35m_cq7MJq70V1zq4xN5pg.gif
// which is significantly negligible compared to that of thread stack size.				local queue	Runnable Goroutines
// Goroutines, CPU Core -> System Thread -> Running Goroutine -> Runnable Goroutines [goroutine1,goroutine2,goroutine(n)....]
// A CPU can have multiple cores, lets focus on one of them for simplicity. These CPU cores are assigned to the threads/process for their execution.
// Since Goroutines are functions which are executed in a thread,
// we say can have multiple Goroutines which are classified into Running Goroutines and Runnable Goroutines.
// A running Goroutine are the ones which are currently running/executing in the thread.
// Whereas Runnable Goroutines are the ones waiting in the local queue to get executed in a thread.
/*
	if a running Goroutine blocks a thread by waiting for I/O,
	a new thread is created and is assigned with a CPU core + Runnable Goroutine that was waiting
	in the local queue, thus continuing execution. When I/O is done, the Goroutine waits in the local queue for its round.
	https://miro.medium.com/v2/resize:fit:720/format:webp/1*wzojTHwqtWDR1CEmTZJKnw.gif
	The above image that the local queue of 1st core is empty. Thus Half of the Goroutines from other local queue are migrated/moved
	to the empty local queue. If every local queue are empty, then Goroutines from global queue are picked.
*/

var wg = &sync.WaitGroup{}

func goroutineFunc2() {
	urls := []string{"https://google.com", "https://medium.com", "https://go.dev"} //domains
	for _, url := range urls {
		wg.Add(1)
		go GetRequest(url)
	}

	wg.Wait()
}

func GetRequest(url string) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Status code for %s is %s\n", url, response.Status)

	wg.Done()
}
