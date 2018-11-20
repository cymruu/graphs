package graphs

import (
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
	size := math.Sqrt(float64(len(adjm)))
	g := CreateEmptyGraph(uint32(size))
	g.AdjacencyMatrix = adjm
	for i := 0; i < int(size); i++ {
		g.Vertices[i] = &Vertex{label: string(rune('A' + i))}
		g.Vertices[i].pos = &positions[i]
	}
	return g
}
func (g *Graph) IsEulers() bool {
	//Euler stwierdził, że aby możliwe było zbudowanie takiej ścieżki, liczba wierzchołków nieparzystego stopnia musi wynosić 0 lub 2.

	return false
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
