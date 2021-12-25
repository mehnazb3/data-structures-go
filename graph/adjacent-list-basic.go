package main

import (
  "fmt"
)

type Graph struct {
	Vertices []*Vertix
}

type Vertix struct {
  Key int
  Edges []*Vertix
}

func (g *Graph) AddVertix(data int) *Vertix {
  if getVertixByKey(g.Vertices, data) != nil {
  	return nil
  }
  newVertix := &Vertix{
  	Key: data,
  }
  g.Vertices = append(g.Vertices, newVertix )
  return newVertix
}

func getVertixByKey(list []*Vertix, data int) *Vertix {
  for _, vertix := range list {
  	if vertix.Key == data {
  		return vertix
  	}
  }
  return nil
}

func (g *Graph) AddEdge(from int, to int, bidirectional bool) {
  fromVertix := getVertixByKey(g.Vertices, from)
  toVertix := getVertixByKey(g.Vertices, to)
  if fromVertix == nil || toVertix == nil {
  	return 
  }

  fromVertix.Edges = append(fromVertix.Edges, toVertix)
  if bidirectional {
  	toVertix.Edges = append(toVertix.Edges, fromVertix)
  }
}

func (g *Graph) Print() {
  for _, vertix := range g.Vertices {
  	fmt.Printf("Vertix %v : \n", vertix.Key)
  	for _, edge := range vertix.Edges {
      fmt.Printf("Edges: %v \n", edge)
  	}
  }
}

func main() {
	var myGraph Graph
	myGraph.AddVertix(1)
	myGraph.AddVertix(2)
	myGraph.AddVertix(2)
	myGraph.AddVertix(3)
	myGraph.AddVertix(4)
	myGraph.AddVertix(5)
	myGraph.AddEdge(2, 3, false)
	myGraph.AddEdge(4, 1, false)
	myGraph.AddEdge(5, 3, true)
	fmt.Println("My Graph: ", myGraph)
	myGraph.Print()
}