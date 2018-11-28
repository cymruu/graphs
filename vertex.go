package graphs

import (
	"fmt"
	"image"
)

type Vertex struct {
	label  string
	degree uint32

	pos *image.Point
}

func (v *Vertex) printDegree() {
	fmt.Printf("%s degree: %d\n", v.label, v.degree)
}
