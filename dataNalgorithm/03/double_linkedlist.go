package main

import "fmt"

type Node struct {
	property     int
	nextNode     *Node
	previousNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (l *LinkedList) NodeBetweenValues(firstProperty, secondProperty int) *Node {
	for node := l.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.property == firstProperty && node.nextNode.property == secondProperty {
				return node
			}
		}
	}
	return nil
}

func (l *LinkedList) AddToHead(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	if l.headNode != nil {
		node.nextNode = l.headNode
		l.headNode.previousNode = node
	}
	l.headNode = node
}

func (l *LinkedList) NodeWithValue(nodeProperty int) *Node {
	for node := l.headNode; node != nil; node = node.nextNode {
		if node.property == nodeProperty {
			return node
		}
	}
	return nil
}

func (l *LinkedList) AddToAfter(nodeProperty, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	nodeWith := l.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith
		if nodeWith.nextNode != nil {
			nodeWith.nextNode.previousNode = node
		}
		nodeWith.nextNode = node
	}
}

func (l *LinkedList) LastNode() *Node {
	for node := l.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			return node
		}
	}
	return nil
}

func (l *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	lastNode := l.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	} else {
		l.headNode = node
	}
}

func main() {
	var linkedList LinkedList
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(5)
	linkedList.AddToAfter(1, 7)
	fmt.Println(linkedList.headNode.property)
	node := linkedList.NodeBetweenValues(1, 5)
	if node != nil {
		fmt.Println(node.property)
	} else {
		fmt.Println("Node not found")
	}
}
