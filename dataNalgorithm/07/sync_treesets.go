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
		insertTreeNode(tree.rootNode, treeNode)
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

func (tree *BinarySearchTree) RemoveNode(key int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	removeNode(tree.rootNode, key)
}

func removeNode(treeNode *TreeNode, key int) *TreeNode {
	if treeNode == nil {
		return nil
	}

	if key < treeNode.key {
		treeNode.leftNode = removeNode(treeNode.leftNode, key)
		return treeNode
	}
	if key > treeNode.key {
		treeNode.rightNode = removeNode(treeNode.rightNode, key)
		return treeNode
	}

	if treeNode.leftNode == nil && treeNode.rightNode == nil {
		treeNode = nil
		return nil
	}

	if treeNode.leftNode == nil {
		treeNode = treeNode.rightNode
		return treeNode
	}

	if treeNode.rightNode == nil {
		treeNode = treeNode.leftNode
		return treeNode
	}
	var leftmostrightNode *TreeNode
	leftmostrightNode = treeNode.rightNode
	for {
		if leftmostrightNode != nil && leftmostrightNode.leftNode != nil {
			leftmostrightNode = leftmostrightNode.leftNode
		} else {
			break
		}
	}
	treeNode.key, treeNode.value = leftmostrightNode.key, leftmostrightNode.value
	treeNode.rightNode = removeNode(treeNode.rightNode, treeNode.key)
	return treeNode
}

func (treeset *TreeSet) InsertTreeNode(treeNodes ...TreeNode) {
	var treeNode TreeNode
	for _, treeNode = range treeNodes {
		treeset.bst.InsertElement(treeNode.key, treeNode.value)
	}
}

func (tree *BinarySearchTree) String() {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	fmt.Println("----------------------------------------")
	stringify(tree.rootNode, 0)
	fmt.Println("----------------------------------------")

}

func stringify(treeNode *TreeNode, level int) {
	if treeNode != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += " "
		}
		format += "***>"
		level++
		stringify(treeNode.leftNode, level)
		fmt.Printf(format+"%d\n", treeNode.key)
		stringify(treeNode.rightNode, level)
	}
}

func main() {
	var treeset *TreeSet = &TreeSet{}
	treeset.bst = &BinarySearchTree{}
	var node1 TreeNode = TreeNode{8, 8, nil, nil}
	var node2 TreeNode = TreeNode{3, 3, nil, nil}
	var node3 TreeNode = TreeNode{10, 10, nil, nil}
	var node4 TreeNode = TreeNode{1, 1, nil, nil}
	var node5 TreeNode = TreeNode{6, 6, nil, nil}
	treeset.InsertTreeNode(node1, node2, node3, node4, node5)
	treeset.bst.String()
}
