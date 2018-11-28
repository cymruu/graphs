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
func (g *Graph) IsEulerian() bool {
	//An Eulerian graph is a graph containing an Eulerian cycle.
	// http://mathworld.wolfram.com/EulerianGraph.html
	//Theorem 1: A graph G=(V(G),E(G)) is Eulerian if and only if each vertex has an even degree.
	//(http://mathonline.wikidot.com/eulerian-graphs-and-semi-eulerian-graphs)
	for _, vertex := range g.Vertices {
		vertex.printDegree()
		if vertex.degree%2 != 0 {
			return false
		}
	}
	return true
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
