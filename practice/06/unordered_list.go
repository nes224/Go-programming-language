package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}

type UnOrderedList struct {
	headNode *Node
}

func (UnOrderedList *UnOrderedList) AddToHead(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	if UnOrderedList.headNode != nil {
		node.nextNode = UnOrderedList.headNode
	}
	UnOrderedList.headNode = node
}

func (UnOrderedList *UnOrderedList) InterateList() {
	var node *Node
	for node = UnOrderedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func main() {
	var unOrderedList UnOrderedList
	unOrderedList = UnOrderedList{}
	buildOrder := []int{1,3,5,7}
	for i := range buildOrder {
		unOrderedList.AddToHead(buildOrder[i])
	}
	unOrderedList.InterateList()

}
