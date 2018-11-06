package graphs

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
	"time"
)

var imageSize = 640
var vertexSize = 10

func (g *Graph) placeVerticestInRandomPlaces() {
	for _, vertex := range g.Vertices {
		x, y := random.Intn(imageSize-vertexSize), random.Intn(imageSize-vertexSize)
		vertex.pos = image.Point{x, y}
	}
}
func (g *Graph) ToImage() {
	img := image.NewRGBA(image.Rect(0, 0, imageSize, imageSize))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{0, 100, 0, 255}}, image.ZP, draw.Src)
	g.placeVerticestInRandomPlaces()
	for _, vertex := range g.Vertices {
		point := image.Rect(vertex.pos.X, vertex.pos.Y, vertex.pos.X+vertexSize, vertex.pos.Y+vertexSize)
		vertexColor := image.Uniform{color.RGBA{0, 255, 0, 255}}
		draw.Draw(img, point, &vertexColor, image.ZP, draw.Src)
		fmt.Printf("%s pos: [%d,%d]\n", vertex.label, vertex.pos.X, vertex.pos.Y)
		/*
			int array[width * height];

			 int SetElement(int row, int col, int value)
			 {
			    array[width * row + col] = value;
			 }
		*/
		for i := 0; i < len(g.AdjacencyMatrix); i++ {
			if g.AdjacencyMatrix[i] {
				// start := g.Vertices[i%graphSize].pos
				// end := image.Point{0, 0}
				// drawLine(img, start, end, color.RGBA{0, 0, 0, 0})
			}
		}
		out, _ := os.Create(fmt.Sprintf("./out/%d.png", time.Now().Unix()))
		png.Encode(out, img)
	}
}
func drawLine(img draw.Image, start, end image.Point,
	fill color.Color) {
	x0, x1 := start.X, end.X
	y0, y1 := start.Y, end.Y
	Δx := math.Abs(float64(x1 - x0))
	Δy := math.Abs(float64(y1 - y0))
	if Δx >= Δy { // shallow slope
		if x0 > x1 {
			x0, y0, x1, y1 = x1, y1, x0, y0
		}
		y := y0
		yStep := 1
		if y0 > y1 {
			yStep = -1
		}
		remainder := float64(int(Δx/2)) - Δx
		for x := x0; x <= x1; x++ {
			img.Set(x, y, fill)
			remainder += Δy
			if remainder >= 0.0 {
				remainder -= Δx
				y += yStep
			}
		}
	} else { // steep slope
		if y0 > y1 {
			x0, y0, x1, y1 = x1, y1, x0, y0
		}
		x := x0
		xStep := 1
		if x0 > x1 {
			xStep = -1
		}
		remainder := float64(int(Δy/2)) - Δy
		for y := y0; y <= y1; y++ {
			img.Set(x, y, fill)
			remainder += Δx
			if remainder >= 0.0 {
				remainder -= Δy
				x += xStep
			}
		}
	}
}
