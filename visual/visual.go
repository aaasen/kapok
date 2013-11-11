package visual

import (
	"fmt"
	"io"
	"math"

	"github.com/aaasen/kapok/graph"
	svg "github.com/ajstarks/svgo"
)

const width = 2048
const height = width

func Visualise(g *graph.Graph, writer io.Writer) *svg.SVG {
	canvas := svg.New(writer)
	canvas.Start(width, height)

	positions := make(map[string]*Vector)
	levels := make(map[int][]*graph.Node)

	for node, _ := range g.Nodes {
		mag := len(g.Nodes[node])

		if levels[mag] == nil {
			levels[mag] = make([]*graph.Node, 0)
		}

		levels[mag] = append(levels[mag], node)
	}

	for mag, nodes := range levels {
		angle := 0.0

		for _, node := range nodes {
			rad := (width / 2.5) / float64(mag+1)

			rotationCorrection := 0.0

			if angle > math.Pi/2.0 && angle < 3*(math.Pi/2) {
				rotationCorrection = math.Pi
			}

			positions[node.Name] = &Vector{
				X:   int((width / 2) + math.Cos(angle)*rad),
				Y:   int((height / 2) + math.Sin(angle)*rad),
				Rot: (angle + rotationCorrection) * (180 / math.Pi),
				Rad: rad,
			}

			angle += (math.Pi * 2) / float64(len(nodes))
		}
	}

	for node, _ := range g.Nodes {
		for neighbor, _ := range g.Nodes[node] {
			vector := positions[node.Name]
			neighborVec := positions[neighbor.Name]

			canvas.Line(
				vector.X, vector.Y,
				neighborVec.X, neighborVec.Y,
				"stroke:#e74c3c;stroke-width:0.5")
		}
	}

	for node, vector := range positions {
		canvas.Circle(
			vector.X,
			vector.Y,
			1,
			"fill:rgba(0, 0, 0, 0.2)")

		canvas.Text(
			0,
			0,
			node,
			fmt.Sprintf(`style="text-anchor:middle;font-size:%vpx;fill:rbga(0, 0, 0, 1);"`,
				int(64.0/(vector.Rad))+12),
			fmt.Sprintf(`transform="rotate(%v, %v, %v) translate(%v, %v)"`,
				vector.Rot,
				vector.X, vector.Y,
				vector.X, vector.Y))
	}

	canvas.End()

	return canvas
}
