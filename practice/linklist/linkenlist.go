package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) prepapreList(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l LinkedList) printData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *LinkedList) deleteValue(k int) {
	if l.length == 0 {
		return
	}

	if l.head.data == k {
		l.head = l.head.next
		l.length--
		return
	}

	previousDeleteValues := l.head
	for previousDeleteValues.next.data != k {
		if previousDeleteValues.next.next == nil {
			return
		}
		previousDeleteValues = previousDeleteValues.next
	}
	previousDeleteValues.next = previousDeleteValues.next.next
	l.length--
}

func LinkList() {
	myList := LinkedList{}
	node1 := &Node{data: 20}
	node2 := &Node{data: 11}
	node3 := &Node{data: 9}
	node4 := &Node{data: 15}
	node5 := &Node{data: 17}
	node6 := &Node{data: 3}
	myList.prepapreList(node1)
	myList.prepapreList(node2)
	myList.prepapreList(node3)
	myList.prepapreList(node4)
	myList.prepapreList(node5)
	myList.prepapreList(node6)

	myList.printData()
	myList.deleteValue(3)
	myList.printData()
	myList.deleteValue(11)
	myList.printData()

}
