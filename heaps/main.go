package main

import "fmt"

// Normal Queues -> First In First Out
// Priority Queues -> Highest priority served first
// Max Heap, Min Heap
// parent index -> -- parent index * 2 + 1 = left child index
// parent index -> -- parent index * 2 + 2 = right child index

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	// when the array is empty
	l := len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}
	// take out the last index and put it in the root
	h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.maxHeapifyDown(0)

	return extracted
}


// maxHeapifyUp will heapify from bottom top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown will heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l,r := left(index), right(index)
	childToCompare := 0
	// lopp while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}
		// compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else { // if the vlaue of the current index is actually larger than the larger child that means that it has found its right place.
			return 
		}
	}
}

// get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}
	fmt.Printf("%v\n", m)
	buildHeap := []int{10, 20, 30, 5,7,9,11,13,15,17}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Printf("%v\n", m)
	}
	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
