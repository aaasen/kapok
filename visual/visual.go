package visual

import (
	"io"
	"math/rand"

	"github.com/aaasen/kapok/graph"
	svg "github.com/ajstarks/svgo"
)

const width = 1024
const height = 1024

type positions struct {
	Positions map[string]*Vector
}

func newPositions() *positions {
	return &positions{
		Positions: make(map[string]*Vector),
	}
}

func (self *positions) SafeGet(name string) *Vector {
	vector := self.Positions[name]

	if vector == nil {
		vector = &Vector{
			X: rand.Intn(width),
			Y: rand.Intn(height),
		}
		self.Positions[name] = vector
	}

	return vector
}

func Visualise(g *graph.Graph, writer io.Writer) *svg.SVG {
	canvas := svg.New(writer)
	canvas.Start(width, height)

	positionsA := newPositions()

	for node, _ := range g.Nodes {
		nodePos := positionsA.SafeGet(node.Name)

		canvas.Circle(
			nodePos.X,
			nodePos.Y,
			1,
			"fill:black")

		for neighbor, _ := range g.Nodes[node] {
			neighborPos := positionsA.SafeGet(neighbor.Name)

			canvas.Line(
				nodePos.X, nodePos.Y,
				neighborPos.X, neighborPos.Y,
				"stroke:rgba(0, 0, 0, 0.2);stroke-width:0.5")
		}
	}

	canvas.End()

	return canvas
}
