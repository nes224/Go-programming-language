package main

// we put key words something into the hash function and it's going to spit out 82
// each letter of the word RANDY R = 82, A = 65, N = 78, D = 68, Y = 89 will be converted into it ansi code
// this just means that we're going to turn each letter into a number and then we sum it up and divide it with 100, 82+65+78+68+89 = 382 / 100 = 82
// so why is it 100 because that is the size of the array that we want to store the name into in.

// Collisions
// in the hash function might be a chance that two names can have the same hash 
// INSERT RANDY -> 4, INSERT ERIC ->> 4 hash function they will be giving out the same hash code
// if we had stored Randy in the array and then after that we try to store Eric that index number if already taken by Randy so we can't store Eric
// To handling these kinds of situation is called Collision handling
// Collision handling would let us find a way to store both Randy and Eric in the array
// Collsion handling -> 1. open addressing, 2. separate chaining
// Open adddressing if the spot for Randy is already taken we just store Eric in the next index and when we later on try to search Eroc
// We first go to its original postion where the hash functions tells us to goif it's not there then we'll look into the next index
// Separate chaining -> Storing multiple names in One index by using linked lists (key,bucket)

import "fmt"

const Arraysize = 7
// HashTable structure
type HashTable struct {
	array [Arraysize]*bucket
}
// bucket structure
type bucket struct {
	head *bucketNode
}
// bucketNode structure
type bucketNode struct {
	key string
	next *bucketNode
}

// Hash table
// Insert
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}
// Search
func (h HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}
// Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// Bucket
// insert will take in a key, create a node with the key and insert the node in the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("Key already exists: %s", k)
	}
}
// search
func (b bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}
// delete
func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// hashfunction
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % Arraysize
}
// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	hashTable := Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}
	for _, v := range list {
		hashTable.Insert(v)
	}
	hashTable.Delete("STAN")
	fmt.Println("STAN", hashTable.Search("STAN"))
	fmt.Println("KENNY", hashTable.Search("KENNY"))
}