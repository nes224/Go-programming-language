package main

const (
	messagePassStart = iota
	messageTicketStart
	messagePassEnd
	messageTicketEnd
)

type QueueLis struct {
	waitPass    int
	waitTicket  int
	playPass    bool
	playTicket  bool
	queuePass   chan int
	queueTicket chan int
	message     chan int
}

func SyncQueue() {
	// var queue *QueueLis = &QueueLis{}
}
