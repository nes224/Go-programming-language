package main

import (
	"fmt"
	"sync"
)

type TreeNode struct {
	key       int
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

type BinarySearchTree struct {
	rootNode *TreeNode
	lock     sync.RWMutex
}

type TreeSet struct {
	bst *BinarySearchTree
}

func (tree *BinarySearchTree) InsertElement(key int, value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	var treeNode *TreeNode
	treeNode = &TreeNode{key, value, nil, nil}
	if tree.rootNode == nil {
		tree.rootNode = treeNode
	} else {

	}
}

func insertTreeNode(rootNode *TreeNode, newTreeNode *TreeNode) {
	if newTreeNode.key < rootNode.key {
		if rootNode.leftNode == nil {
			rootNode.leftNode = newTreeNode
		} else {
			insertTreeNode(rootNode.leftNode, newTreeNode)
		}
	} else {
		if rootNode.rightNode == nil {
			rootNode.rightNode = newTreeNode
		} else {
			insertTreeNode(rootNode.rightNode, newTreeNode)
		}
	}
}

func (treeset *TreeSet) SplitKeyValues(treeNodes ...TreeNode) {
	var treeNode TreeNode
	for _, treeNode = range treeNodes {
		treeset.bst.InsertElement(treeNode.key, treeNode.value)
	}
}

func BinarySearchTreeFn() {
	var treeset *TreeSet = &TreeSet{}
	treeset.bst = &BinarySearchTree{}
	node1 := TreeNode{8, 8, nil, nil}
	node2 := TreeNode{3, 3, nil, nil}
	node3 := TreeNode{10, 10, nil, nil}
	node4 := TreeNode{1, 1, nil, nil}
	node5 := TreeNode{6, 6, nil, nil}
	treeset.SplitKeyValues(node1, node2, node3, node4, node5)

	fmt.Println(treeset)
}
