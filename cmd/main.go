package main

import (
	"fmt"

	"github.com/cymruu/graphs"
)

func main() {
	graph := graphs.CreateRandomGraph(6, .5)

	fmt.Print(graph)
	graph.RandomizePoints()
	graph.ToImage()
}
