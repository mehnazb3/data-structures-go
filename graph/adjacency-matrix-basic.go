package main

import (
	"fmt"
)

type Graph struct {
	Vertices []int
	Edges    [][]int
}

func (g *Graph) AddVerticex(data int) {
	if getIndex(g.Vertices, data) >= 0 {
		return
	}
	g.Vertices = append(g.Vertices, data)
	size := len(g.Vertices)
	existingData := g.Edges
	g.Edges = make([][]int, size)
	for i := 0; i < len(g.Vertices); i++ {
		g.Edges[i] = make([]int, size)
		for j := 0; j < len(g.Vertices); j++ {
			if len(existingData)-1 > i && len(existingData)-1 > j {
				g.Edges[i][j] = existingData[i][j]
			}
		}
	}
}

func getIndex(list []int, item int) int {
	for i, v := range list {
		if v == item {
			return i
		}
	}
	return -1
}

func (g *Graph) AddEdge(from, to int, bidirectional bool) {
	fromIndex := getIndex(g.Vertices, from)
	toIndex := getIndex(g.Vertices, to)
	if fromIndex >= 0 && toIndex >= 0 {
		g.Edges[fromIndex][toIndex] = 1
		if bidirectional {
			g.Edges[toIndex][fromIndex] = 1
		}
	}
}

func main() {
	var myGraph Graph
	myGraph.AddVerticex(1)
	myGraph.AddVerticex(2)
	myGraph.AddVerticex(3)
	myGraph.AddVerticex(3) // No change
	myGraph.AddVerticex(4)
	myGraph.AddEdge(1, 3, false)
	myGraph.AddVerticex(5)
	myGraph.AddEdge(5, 2, true)
	myGraph.AddEdge(6, 2, true) // No Change
	myGraph.AddEdge(1, 1, true) // Loop

	fmt.Println("My Graph: ", myGraph)
}
