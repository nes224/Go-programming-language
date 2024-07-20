package main

/*
how to wait for other goroutines to finish instead of exiting on the main goroutine
func main() {
	done := make(chan struct{}) // unbuffered channel is that it is blocking until it receives a value from the channel.
	go func(){
		// ...
		done <- struct{}{} // signal the main goroutine
	}())
	// ...
	<- done // wait for background goroutine to finish
}
*/

/*
unbuffered vs. buffered channel
potential issue with the using the unbuffered 
channel below?
func mirroredQuery() string {
	responses := make(chan string)
	/* the fastest of the three goroutines will return a response and the function will exit.
	go func() { responses <- "response 1"} ()
	go func() { responses <- "response 2"} ()
	go func() { responses <- "response 3"} ()
	/* the two slower goroutine would be stuck trying to send their responses on a channel from which no goroutine will eve receive
	return <- responses // return the quickest response

	goroutine leak, it is important to ensure that go routines terminate themselves.
	too many arguments to return
	have (string)
	want ()
}
*/

/*
func mirroredQuery() string {
	/* assign a capacity of 3 to the channel
	responses := make(chan string, 3) // bufferen channel
	go func() { responses <- "response 1" } ()
	go func() { responses <- "response 2" } ()
	go func() { responses <- "response 3" } ()
	/* the channel would hold up to three string values and
	block until a space is made available by another
	goroutine's receive
	return <- responses // return the quickest response

	// do not use the buffered channel as a queue within a single goroutine
}
*/


/*
unbuffered vs. buffered channel
unbuffered channel
 - stronger synchronization guarantees
 - risk of goroutine leaks

buffered channel
 - reduced risk of goroutine leaks
 - failure to allocate sufficient buffer capacity can lead to deadlock
*/

/*
func makeThumbnails(filesnames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1) // incrementing goroutines
		go func(f string) {
			defer wg.Done() // decrementing goroutines
			sizes <- size
		}(f)
	}

	go func() {
		wg.Wait() // ensures that we wait until all the go routines have executed and Done() has been called on each of the goroutines that we added.
		close(sizes)
	}

	var total int64
	for size := range sizes {
		total += size
	}

	return total
}

*/