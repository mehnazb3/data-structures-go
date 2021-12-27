package main

import "fmt"

type Graph struct {
	Bidirectional bool
	Vertices      map[int]*Vertix
}

type Vertix struct {
	Key       int
	Adjacents map[int]*Vertix
}

type LinkedList struct {
	Head *Node
}

type Node struct {
	Key  int
	Next *Node
}

func NewVertix(key int) *Vertix {
	return &Vertix{
		Key:       key,
		Adjacents: make(map[int]*Vertix),
	}
}

func (g *Graph) AddVertix(vertix *Vertix) *Vertix {
	if _, ok := g.Vertices[vertix.Key]; ok {
		return nil
	}
	g.Vertices[vertix.Key] = vertix
	return vertix
}

func (g *Graph) findVertixBykey(key int) *Vertix {
	return g.Vertices[key]
}

func (g *Graph) AddEdge(from, to int) {
	fromVertix := g.findVertixBykey(from)
	toVertix := g.findVertixBykey(to)
	if fromVertix != nil && toVertix != nil {
		fromVertix.Adjacents[to] = toVertix
		if g.Bidirectional {
			toVertix.Adjacents[from] = fromVertix
		}
	}
	return
}

func contains(list []int, key int) bool {
	for _, item := range list {
		if item == key {
			return true
		}
	}
	return false
}

func (g *Graph) DFS(startNode *Vertix, visitedVertix func(int)) {
	visited := make(map[int]bool)
	visited[startNode.Key] = true
	visitedVertix(startNode.Key)
	for _, adjacent := range startNode.Adjacents {
		if _, ok := visited[adjacent.Key]; ok {
			continue
		}
		g.DFS(adjacent, visitedVertix)
	}
}

func (l *LinkedList) enqueue(key int) {
	currentNode := l.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = &Node{
		Key: key,
	}
}

func (g *Graph) BFS(startNode *Vertix, visitedVertixBFS func(int)) {
	visited := make(map[int]bool)

	queue := &LinkedList{
		Head: &Node{
			Key: startNode.Key,
		},
	}
	node := queue.Head
	for node != nil {
		visited[node.Key] = true
		visitedVertixBFS(node.Key)
		for _, adjacent := range g.Vertices[node.Key].Adjacents {
			if _, ok := visited[adjacent.Key]; ok {
				continue
			}
			queue.enqueue(adjacent.Key)
		}
		node = node.Next
	}
}

func main() {
	myGraph := Graph{
		Vertices: make(map[int]*Vertix),
	}
	// Add Vertices
	startNode := NewVertix(1)
	myGraph.AddVertix(startNode)
	myGraph.AddVertix(NewVertix(2))
	myGraph.AddVertix(NewVertix(2))
	myGraph.AddVertix(NewVertix(3))
	myGraph.AddVertix(NewVertix(4))
	myGraph.AddVertix(NewVertix(5))
	myGraph.AddVertix(NewVertix(6))
	myGraph.AddVertix(NewVertix(7))
	// Add Edges
	myGraph.AddEdge(1, 2)
	myGraph.AddEdge(1, 3)
	myGraph.AddEdge(1, 6)
	myGraph.AddEdge(6, 5)
	myGraph.AddEdge(3, 5)
	myGraph.AddEdge(2, 4)
	myGraph.AddEdge(4, 3)
	myGraph.AddEdge(5, 7)
	// Depth first search
	fmt.Println("DFS traversing")
	visited := []int{}
	visitedVertix := func(visitedKey int) {
		if !contains(visited, visitedKey) {
			visited = append(visited, visitedKey)
			fmt.Println("Visiting : ", visitedKey)
		}
	}
	myGraph.DFS(startNode, visitedVertix)
	fmt.Println("BFS traversing")
	bfsVisited := []int{}
	visitedVertixBFS := func(visitedKey int) {
		if !contains(bfsVisited, visitedKey) {
			bfsVisited = append(bfsVisited, visitedKey)
			fmt.Println("Visiting : ", visitedKey)
		}
	}
	myGraph.BFS(startNode, visitedVertixBFS)

}
