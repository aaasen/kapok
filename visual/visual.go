package visual

import (
	"io"
	"math/rand"

	"github.com/aaasen/kapok/graph"
	svg "github.com/ajstarks/svgo"
)

func Visualise(graph *graph.Graph, writer io.Writer) *svg.SVG {
	width := 1024
	height := 1024
	canvas := svg.New(writer)
	canvas.Start(width, height)

	for k, _ := range graph.Nodes {
		canvas.Text(int(rand.Float32()*float32(width)), int(rand.Float32()*float32(height)), k.Name, "text-anchor:middle;font-size:12px;fill:black")
	}

	canvas.End()

	return canvas
}
