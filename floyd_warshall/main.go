package main

import (
	"fmt"
	"math"
)

const inf = math.MaxInt

type Graph struct {
	vertices [][]int
}

func NewGraph(size int) *Graph {
	vertices := make([][]int, size)
	for i := range vertices {
		vertices[i] = make([]int, size)
		for j := range vertices[i] {
			if i == j {
				vertices[i][j] = 0
			} else {
				vertices[i][j] = inf
			}
		}
	}

	return &Graph{vertices: vertices}
}

func (g *Graph) AddEdge(u, v int, weight int) {
	g.vertices[u][v] = weight
}

func (g *Graph) FloydWarshall() [][]int {
	dist := make([][]int, len(g.vertices))
	for i := range dist {
		dist[i] = make([]int, len(g.vertices))
		copy(dist[i], g.vertices[i])
	}

	for k := range g.vertices {
		for i := range g.vertices {
			for j := range g.vertices {
				if dist[i][k] < inf && dist[k][j] < inf && dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	return dist
}

func main() {
	graph := NewGraph(4)
	graph.AddEdge(0, 1, 5)
	graph.AddEdge(0, 3, 10)
	graph.AddEdge(1, 2, 3)
	graph.AddEdge(2, 3, 1)

	distance := graph.FloydWarshall()
	for _, dist := range distance {
		for _, d := range dist {
			if d == inf {
				fmt.Printf("inf ")
			} else {
				fmt.Printf("%d ", d)
			}
		}
		fmt.Println()
	}
}
