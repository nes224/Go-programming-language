package main

// every single thread consumes ~1MB of memory. Thus if your application creates 1000 threads then it would consume ~1GB, which is huge.
// light weight process. light weight threads.
// Goroutines are managed by Go Runtime. Every single goroutine consume ~2KB, 
// which is significantly negligible compared to that of thread stack size.
// Goroutines, CPU Core -> System Thread -> Running Goroutine
// A CPU can have multiple cores, lets focus on one of them for simplicity. These CPU cores are assigned to the threads/process for their execution.
// Since Goroutines are functions which are executed in a thread,
// we say can have multiple Goroutines which are classified into Running Goroutines and Runnable Goroutines.
// A running Goroutine are the ones which are currently running/executing in the thread.
// Whereas Runnable Goroutines are the ones waiting in the local queue to get executed in a thread.
func goroutineFunc2() {

}