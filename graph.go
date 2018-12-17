package graphs

import (
	"fmt"
	"image"
	"math"
	"math/rand"
	"time"
)

type Graph struct {
	Vertices        []*Vertex
	AdjacencyMatrix []bool
}
type Pos struct {
	X, Y int
}

var sourceUnix = rand.NewSource(time.Now().UnixNano())
var staticSource = rand.NewSource(100023)
var random = rand.New(sourceUnix)

func CreateEmptyGraph(size uint32) *Graph {
	g := &Graph{
		Vertices:        make([]*Vertex, size),
		AdjacencyMatrix: make([]bool, size*size),
	}
	for i := 0; i < len(g.AdjacencyMatrix); i++ {
		g.AdjacencyMatrix[i] = false
	}
	return g
}
func (g *Graph) Size() int {
	return len(g.Vertices)
}
func CreateRandomGraph(numberOfVertices uint32) *Graph {
	g := CreateEmptyGraph(numberOfVertices)
	g.Vertices[0] = &Vertex{label: string(rune('A'))}
	for i := uint32(1); i < numberOfVertices; i++ {
		vertex := &Vertex{label: string(rune('A' + i))}
		g.Vertices[i] = vertex
		randomVertexIndex := random.Intn(cap(g.AdjacencyMatrix))
		g.AdjacencyMatrix[randomVertexIndex] = true
	}
	return g
}
func CreateGraph(adjm []bool, positions []image.Point) *Graph {
	if len(positions) != int(math.Pow(float64(len(adjm)), 2)) {
		//error
	}
	size := math.Sqrt(float64(len(adjm)))
	g := CreateEmptyGraph(uint32(size))
	g.AdjacencyMatrix = adjm
	for i := 0; i < int(size); i++ {
		g.Vertices[i] = &Vertex{label: string(rune('A' + i))}
		g.Vertices[i].pos = &positions[i]
	}
	return g
}
func (g *Graph) CalculateDegrees() bool {
	for i := 0; i < g.Size(); i++ {
		degree := uint32(0)
		for j := 0; j < g.Size(); j++ {
			fmt.Print(g.AdjacencyMatrix[i*g.Size()+j], " ")
			if g.AdjacencyMatrix[i*g.Size()+j] {
				degree++
			}
		}
		g.Vertices[i].degree = degree
	}
	return false
}
func (g *Graph) GetVertexAdjacent(vertex *Vertex) []*Vertex {
	adjacent := make([]*Vertex, 0)
	var index int
	for i := 0; i < g.Size(); i++ {
		if g.Vertices[i] == vertex {
			index = i
		}
	}
	row := index % g.Size()
	for i := 0; i < g.Size(); i++ {
		if g.AdjacencyMatrix[row*g.Size()+i] {
			adjacent = append(adjacent, g.Vertices[i])
		}
	}
	return adjacent
}
func (g *Graph) DFS(vertex *Vertex, visited map[*Vertex]bool) {
	visited[vertex] = true
	adjacent := g.GetVertexAdjacent(vertex)
	for _, vertex := range adjacent {
		if !visited[vertex] {
			g.DFS(vertex, visited)
		}
	}
}

//
func (g *Graph) IsConnected() bool {
	visited := make(map[*Vertex]bool)
	for _, vertex := range g.Vertices {
		visited[vertex] = false
	}
	var start *Vertex
	for _, vertex := range g.Vertices {
		if vertex.degree > 0 {
			start = vertex //select first Vertex with edge to start IsConnected Algorithm from
		}
	}
	if start == nil {
		return true //TODO: A graph with just one vertex is connected. An edgeless graph with two or more vertices is disconnected.
	}
	g.DFS(start, visited)
	//check if all non-zero degree vertices are visited
	for _, vertex := range g.Vertices {
		if visited[vertex] == false {
			return false
		}
	}
	return true
}

/* The function returns one of the following values
   0 --> If grpah is not Eulerian
   1 --> If graph has an Euler path (Semi-Eulerian)
   2 --> If graph has an Euler Circuit (Eulerian)  */
func (g *Graph) IsEulerian() int {
	//An Eulerian graph is a graph containing an Eulerian cycle.
	// http://mathworld.wolfram.com/EulerianGraph.html
	//Theorem 1: A graph G=(V(G),E(G)) is Eulerian if and only if each vertex has an even degree.
	//(http://mathonline.wikidot.com/eulerian-graphs-and-semi-eulerian-graphs)
	//Algorithm:
	//https://www.geeksforgeeks.org/fleurys-algorithm-for-printing-eulerian-path/
	if !g.IsConnected() {
		return 0
	}
	odd := 0
	for _, vertex := range g.Vertices {
		if vertex.degree%2 != 0 {
			odd++
		}
	}
	if odd == 0 {
		return 2
	} else if odd == 2 {
		return 1
	} else {
		return 0
	}
}
func (g *Graph) Copy() *Graph {
	graph := &Graph{
		Vertices:        make([]*Vertex, len(g.Vertices)),
		AdjacencyMatrix: make([]bool, len(g.AdjacencyMatrix)),
	}
	copy(graph.Vertices, g.Vertices)
	copy(graph.AdjacencyMatrix, g.AdjacencyMatrix)
	return graph
}
