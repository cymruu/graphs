package graphs

import (
	"math/rand"
	"time"
)

type AdjacencyMatrixElem struct {
	*Vertex
	NumberOfEdges float64
}
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

func CreateEmptyGraph(size int) *Graph {
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
func (g *Graph) IndexToPos(i int) Pos {
	return Pos{i % g.Size(), i / g.Size()}
}
func CreateRandomGraph(numberOfVertices int) *Graph {
	g := CreateEmptyGraph(numberOfVertices)
	g.Vertices[0] = &Vertex{label: string(rune('A'))}
	for i := 1; i < numberOfVertices; i++ {
		vertex := &Vertex{label: string(rune('A' + i))}
		g.Vertices[i] = vertex
		randomVertexIndex := random.Intn(i)
		g.AdjacencyMatrix[randomVertexIndex] = true
	}
	return g
}
