package main

import "fmt"

var count int

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}

	}
}

func (n *Node) Search(k int) bool {
	count++
	if n == nil {
		return false
	}
	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}
	return true
}

func main() {
	tree := &Node{Key: 100}
	treeNumber := []int{52, 203, 19, 76, 150, 310, 7, 24, 88, 276}
	for i := 0; i < len(treeNumber); i++ {
		tree.Insert(treeNumber[i])
	}
	fmt.Println(tree.Right)

	fmt.Println(tree.Search(400))

	fmt.Println(tree.Search(7))
	fmt.Printf("%d\n",count)

}
