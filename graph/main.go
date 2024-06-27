package main
import (
	"fmt"
)

// Adjacency List -> Key,Values  node is the edge and director is a vertex
// 1,2,3,4,5
// Adjacency list of vertex 2
// Graph structure
// Adjacency Matrix (x,y)
type Graph struct {
	vertices []*Vertex
}
// Vertex structure
type Vertex struct {
	key int
	adjacent []*Vertex
}

// AddVertex adds a Vertex to the Graph
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key:k})
	}
}
// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v-->%v)",from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent,to){
		err := fmt.Errorf("Existing edge (%v-->%v)",from,to)
		fmt.Println(err.Error())
	}else {
		// add edge
		fromVertex.adjacent = append(fromVertex.adjacent,toVertex)
	}
}

// get Vertex
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

//contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}


// Print will print the adjacent list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices{
		fmt.Printf("\nVertex %v :", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}
}


func main() {
	test := &Graph{}

	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}
	test.AddEdge(1,2)
	test.AddEdge(1,3)
	test.AddEdge(2,3)
	test.AddEdge(2,4)
	test.AddEdge(1,2)
	test.AddEdge(6,2)
	test.AddEdge(3,2)
	test.Print()
}