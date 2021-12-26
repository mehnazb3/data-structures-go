// imagine a table with 3 rows and 4 columns
// Each column has differen color
// Find out max count of same color connected column?
// For below problem, there are max 5 connected column of same color.
// -----------------
// | R | - | B | G |
// | R | B | G | B |
// | - | B | B | B |
// -----------------
package main

import (
	"fmt"
)

type Color int

// Colors are knowm
const (
	Red   Color = iota // 0
	Black              // 1
	White              // 2
	Green              // 3
)

type Graph struct {
	Vertices []*Vertix
}

type Vertix struct {
	Key       int
	BGColor   Color
	Adjacents []*Vertix
}

type ResponseResult struct {
	Responses []Response
	// Keys []int
}

type Response struct {
	BGColor Color
	Count   int
	Keys    []int
}

func (g *Graph) AddVerticex(vertix *Vertix) {
	if FindVertix(g.Vertices, vertix.Key) != nil {
		return
	}
	g.Vertices = append(g.Vertices, vertix)
}

func FindVertix(vertices []*Vertix, key int) *Vertix {
	for _, vertix := range vertices {
		if vertix.Key == key {
			return vertix
		}
	}
	return nil
}

func contains(list []int, key int) bool {
	for _, item := range list {
		if item == key {
			return true
		}
	}
	return false
}

func (g *Graph) AddEdge(from int, to int) {
	fromVertix := FindVertix(g.Vertices, from)
	toVertix := FindVertix(g.Vertices, to)
	if fromVertix != nil && toVertix != nil {
		fromVertix.Adjacents = append(fromVertix.Adjacents, toVertix)
		toVertix.Adjacents = append(toVertix.Adjacents, fromVertix)
	}
}

func findByBEColor(BGColor Color, adjacents []*Vertix, result []*Response) *Response {
	var responseList []*Response
	for _, response := range result {
		for _, adjacent := range adjacents {
			if response.BGColor == BGColor && contains(response.Keys, adjacent.Key) {
				responseList = append(responseList, response)
			}
		}
	}
	// Concatinate if two elements found
	if len(responseList) == 1 {
		return responseList[0]
	} else if len(responseList) > 1 {
		responseList[0].Count = responseList[0].Count + responseList[1].Count
		responseList[0].Keys = append(responseList[0].Keys, responseList[1].Keys...)
		return responseList[0]
	}
	return nil
}

func (g *Graph) FindConnectedVerticesWithcolor() []*Response {
	var result []*Response
	for _, vertix := range g.Vertices {
		existing := findByBEColor(vertix.BGColor, vertix.Adjacents, result)
		if existing != nil {
			existing.Count += 1
			existing.Keys = append(existing.Keys, vertix.Key)
		} else {
			response := Response{
				BGColor: vertix.BGColor,
				Keys:    []int{vertix.Key},
				Count:   1,
			}
			result = append(result, &response)
		}
	}

	return result
}

func main() {
	var myGraph Graph
	myGraph.AddVerticex(&Vertix{
		Key:     1,
		BGColor: 0,
	})
	myGraph.AddVerticex(&Vertix{
		Key:     2,
		BGColor: 2,
	})
	myGraph.AddVerticex(&Vertix{
		Key:     3,
		BGColor: 1,
	})
	myGraph.AddVerticex(&Vertix{
		Key:     4,
		BGColor: 3,
	})
	myGraph.AddVerticex(&Vertix{
		Key:     5,
		BGColor: 0,
	})
	myGraph.AddVerticex(&Vertix{
		Key:     6,
		BGColor: 1,
	})
	myGraph.AddVerticex(&Vertix{
		Key:     7,
		BGColor: 3,
	})
	myGraph.AddVerticex(&Vertix{
		Key:     8,
		BGColor: 1,
	})
	myGraph.AddVerticex(&Vertix{
		Key:     9,
		BGColor: 2,
	})
	myGraph.AddVerticex(&Vertix{
		Key:     10,
		BGColor: 1,
	})
	myGraph.AddVerticex(&Vertix{
		Key:     11,
		BGColor: 1,
	})
	myGraph.AddVerticex(&Vertix{
		Key:     12,
		BGColor: 1,
	})
	myGraph.AddEdge(1, 2)
	myGraph.AddEdge(1, 5)
	myGraph.AddEdge(2, 3)
	myGraph.AddEdge(2, 6)
	myGraph.AddEdge(3, 4)
	myGraph.AddEdge(3, 7)
	myGraph.AddEdge(4, 8)
	myGraph.AddEdge(5, 6)
	myGraph.AddEdge(5, 9)
	myGraph.AddEdge(6, 7)
	myGraph.AddEdge(6, 10)
	myGraph.AddEdge(7, 8)
	myGraph.AddEdge(7, 11)
	myGraph.AddEdge(8, 12)
	myGraph.AddEdge(9, 10)
	myGraph.AddEdge(10, 11)
	myGraph.AddEdge(11, 12)
	for _, vertix := range myGraph.Vertices {
		var neigh []int
		for _, adj := range vertix.Adjacents {
			neigh = append(neigh, adj.Key)
		}
		fmt.Println("myGraph VERTIX", vertix.Key, neigh)
	}
	res := myGraph.FindConnectedVerticesWithcolor()
	for _, data := range res {
		fmt.Println(data)
	}
}
