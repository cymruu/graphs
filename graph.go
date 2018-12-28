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

const (
	NoPath  int = iota
	Path        = iota
	Circuit     = iota
)

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
func CreateRandomGraph(numberOfVertices int, probability float64) *Graph {
	g := CreateEmptyGraph(uint32(numberOfVertices))
	for i := 0; i < numberOfVertices; i++ {
		vertex := &Vertex{label: string(rune('A' + i))}
		g.Vertices[i] = vertex
	}
	for i := 0; i < len(g.Vertices); i++ {
		for j := 0; j < len(g.Vertices); j++ {
			if g.Vertices[i] != g.Vertices[j] {
				if random.Float64() < probability {
					g.AddEdge(g.Vertices[i], g.Vertices[j])
				}
			}
		}
	}
	g.CalculateDegrees()
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
			if g.AdjacencyMatrix[i*g.Size()+j] {
				degree++
			}
		}
		g.Vertices[i].degree = degree
	}
	return false
}

func (g *Graph) getVertexIndex(vertex *Vertex) int {
	for i := 0; i < g.Size(); i++ {
		if g.Vertices[i] == vertex {
			return i
		}
	}
	return -1
}
func (g *Graph) GetVertexAdjacent(vertex *Vertex) []*Vertex {
	adjacent := make([]*Vertex, 0)
	index := g.getVertexIndex(vertex)

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
	if start == nil && len(g.Vertices) == 1 {
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
		return NoPath
	}
	odd := 0
	for _, vertex := range g.Vertices {
		if vertex.degree%2 != 0 {
			odd++
		}
	}
	if odd == 0 {
		return Circuit
	} else if odd == 2 {
		return Path
	} else {
		return NoPath
	}
}
func (g *Graph) RemoveEdge(v, u *Vertex) {
	g.SetEdge(v, u, false)
}
func (g *Graph) AddEdge(v, u *Vertex) {
	g.SetEdge(v, u, true)
}
func (g *Graph) SetEdge(v, u *Vertex, value bool) {
	vIdx := g.getVertexIndex(v)
	uIdx := g.getVertexIndex(u)

	g.AdjacencyMatrix[vIdx*g.Size()+uIdx] = value
	g.AdjacencyMatrix[uIdx*g.Size()+vIdx] = value
	if value == true {
		g.Vertices[uIdx].degree++
		g.Vertices[vIdx].degree++
	} else {
		g.Vertices[uIdx].degree--
		g.Vertices[vIdx].degree--
	}

}
func (g *Graph) isValidEdgeForEulerianCycle(u, v *Vertex) bool {
	//The edge between u, v is valid in one of the following two cases:
	// 1) v is the only adjacent of vertex of u
	if u.degree == 1 {

		return true
	}
	// 2) If there are multiple adjacents, then u-v is not a bridge
	// Do following steps to check if u-v is a bridge
	// 2.a) count of vertices reachable from u
	visited := make(map[*Vertex]bool)
	counts := make([]int, 2)
	g.DFS(u, visited)
	counts[0] = len(visited)
	// 2.b) Remove edge (u, v) and after removing the edge, count
	// vertices reachable from u
	visited = make(map[*Vertex]bool)
	g.RemoveEdge(u, v)
	g.DFS(u, visited)
	counts[1] = len(visited)
	// 2.c) Add the edge back to the graph
	g.AddEdge(u, v)
	return !(counts[0] > counts[1])
}
func PrintPath(path []*Vertex) {
	fmt.Print("Path: ")
	for _, v := range path[0 : len(path)-1] {
		// fmt.Printf("%+v", v)
		fmt.Printf("%s -> ", v.label)
	}
	fmt.Println(path[len(path)-1].label)
}
func (g *Graph) findNextVertexInEulerianPath(start *Vertex, trail *[]*Vertex) { //TODO: index is unnescessary
	adajecent := g.GetVertexAdjacent(start)
	for _, next := range adajecent {
		if g.isValidEdgeForEulerianCycle(start, next) {
			// trail[index] = next
			*trail = append(*trail, next)
			g.RemoveEdge(start, next)
			g.findNextVertexInEulerianPath(next, trail)
			break
		}
	}
}
func (g *Graph) copy() *Graph {
	new_g := &Graph{
		Vertices:        make([]*Vertex, len(g.Vertices)),
		AdjacencyMatrix: make([]bool, len(g.AdjacencyMatrix)),
	}
	copy(new_g.Vertices, g.Vertices)
	copy(new_g.AdjacencyMatrix, g.AdjacencyMatrix)
	return new_g
}
func (g *Graph) FindEulerianPath() ([]*Vertex, int) {
	graphCopy := g.copy()
	var startVertex *Vertex
	gType := graphCopy.IsEulerian()
	switch gType {
	case NoPath:
		return nil, NoPath
	case 1: //eulerian path (2 odd vertices)
		for _, vertex := range graphCopy.Vertices {
			if vertex.degree%2 != 0 {
				startVertex = vertex
				break
			}
		}
	case 2: //eulerian circle (0 odd vertices)
		startVertex = graphCopy.Vertices[0]
	}
	trail := make([]*Vertex, 0)
	trail = append(trail, startVertex)
	graphCopy.findNextVertexInEulerianPath(startVertex, &trail)
	g.CalculateDegrees()
	return trail, gType
}
func (g *Graph) RandomizePoints() {
	for _, vertex := range g.Vertices {
		if vertex.pos == nil {
			vertex.pos = &image.Point{}
		}
		vertex.pos.X = int(random.Int31n(int32(imageSize - vertexSize)))
		vertex.pos.Y = int(random.Int31n(int32(imageSize - vertexSize)))
	}
}
