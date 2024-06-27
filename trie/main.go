package main

import "fmt"

//				  g -> o -> n
//				/			  g
// 		a -> r -> 			/
//    /			\ a -> g -> o 			   o
// root e->r->a->g->o->n	\ r -> n     / 	 a -> n -> o
//	  \				/ ------> o -> r -> e	/
//		o -> r -> e						 \ g
//											\ o -> n
// [a|0xc0000a6050] | [b|<nil>] | [c|<nil>]

// AlphabetSize is the number of possible characters in the trie
const AlphabetSize = 26

// Node represents each node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie represent a trie and has a pointer t othe root node
type Trie struct {
	root *Node
}

// InitTrie will create new Trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert will take in a  word and add it to the trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a' // a = 97, "b" - "a" = 1
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search will take in a word and RETURN true is that word is included in the trie.
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0 ;i<wordLength;i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}

func main() {
	myTrie := InitTrie()

	toAdd := []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}

	for _, v := range toAdd {
		myTrie.Insert(v)
	}
	fmt.Printf("%v\n", myTrie.root)
	fmt.Println(myTrie.Search("aragorn"))
	fmt.Println(myTrie.Search("go"))
}
