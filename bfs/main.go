package main

import (
	"container/list"
	"fmt"
)

type Graph struct {
	vertexMap map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		vertexMap: make(map[int][]int),
	}
}

func (g *Graph) AddEdge(u, v int) {
	if u != v { // prevent loop
		g.vertexMap[u] = append(g.vertexMap[u], v)
		g.vertexMap[v] = append(g.vertexMap[v], u)
	}
}

func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := list.New()
	queue.PushBack(start)
	visited[start] = true

	for queue.Len() > 0 {
		vertex := queue.Remove(queue.Front()).(int)
		fmt.Printf("%d ", vertex)
		for _, neighbor := range g.vertexMap[vertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.PushBack(neighbor)
			}
		}
	}
}

func main() {
	graph := NewGraph()
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 0)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)

	graph.BFS(4)
}
