package main

import (
	"fmt"
	"time"
)

const (
	messagePassStart = iota
	messageTicketStart
	messagePassEnd
	messageTicketEnd
)

type Queue struct {
	waitPass int
	waitTicket int
	playPass bool
	playTicket bool
	queuePass chan int
	queueTicket chan int
	message chan int
}

func (queue *Queue) New() {
	queue.message = make(chan int)
	queue.queuePass = make(chan int)
	queue.queueTicket = make(chan int)
	go func() {
		var message int
		for {
			select {
			case message = <-queue.message:
				switch message {
				case messagePassStart:
					queue.waitPass++
				case messagePassEnd:
					queue.playPass = false
				case messageTicketStart:
					queue.waitTicket++
				case messageTicketEnd:
					queue.playTicket = false
				}
				if queue.waitPass > 0 && queue.waitTicket > 0 && !queue.playPass && !queue.playTicket {
					queue.playPass = true
					queue.playTicket = true
					queue.waitTicket--
					queue.waitPass--
					queue.queuePass <- 1
					queue.queueTicket <- 1
				}
			}
		}
	}()
}

// method starts the issuing of a ticket for passengers standing in a queue.
func (queue *Queue) StartTicketIssue() {
	queue.message <- messageTicketStart
	<-queue.queueTicket
}

// method finishes the issuing of a ticket to a passenger standing in the queue.
func (queue *Queue) EndTicketIssue() {
	queue.message <- messageTicketEnd
}

func ticketIssue(queue *Queue) {
	for {
		time.Sleep(10 * time.Second)
		queue.StartTicketIssue()
		fmt.Println("Ticket Issue starts")
		time.Sleep(2 * time.Second)
		fmt.Println("Ticket Issue ends")
		queue.EndTicketIssue()
	}
}

// StartPass method starts the passenger queue moving toward the ticket counter.
// Passengers are moved into the queue.
func (queue *Queue) StartPass() {
	queue.message <- messagePassStart
	<-queue.queuePass
}

func (queue *Queue) EndPass() {
	queue.message <- messagePassEnd
}

func passenger(queue *Queue) {
	for {
		// starting the processing
		time.Sleep(10 * time.Second)
		queue.StartPass()
		fmt.Println("Passenger starts")
		time.Sleep(2 *time.Second)
		fmt.Println("Passenger ends")
		queue.EndPass()
	}
}

func main() {
	var queue *Queue = &Queue{}
	queue.New()
	fmt.Println(queue)
	var i int
	for i = 0; i<10; i++ {
		go passenger(queue)
	}
	var j int
	for j = 0; j < 5; j++ {
		go ticketIssue(queue)
	}
	select{}
}
