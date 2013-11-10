package visual

import (
	"fmt"
	"io"
	"math"
	"math/rand"

	"github.com/aaasen/kapok/graph"
	svg "github.com/ajstarks/svgo"
)

const width = 1024
const height = 1024

type positions struct {
	Positions map[string]*Vector
	Graph     *graph.Graph
}

func newPositions(graph *graph.Graph) *positions {
	return &positions{
		Positions: make(map[string]*Vector),
		Graph:     graph,
	}
}

func (self *positions) SafeGet(node *graph.Node) *Vector {
	vector := self.Positions[node.Name]

	if vector == nil {
		angle := rand.Float64() * math.Pi * 2
		r := (width / 2.5) / float64(len(self.Graph.Nodes[node])+1)

		rotationCorrection := 0.0

		if angle > math.Pi/2.0 && angle < 3*(math.Pi/2) {
			rotationCorrection = math.Pi
		}

		vector = &Vector{
			X:   int((width / 2) + math.Cos(angle)*r),
			Y:   int((height / 2) + math.Sin(angle)*r),
			Rot: (angle + rotationCorrection) * (180 / math.Pi),
			Rad: r,
		}

		self.Positions[node.Name] = vector
	}

	return vector
}

func Visualise(g *graph.Graph, writer io.Writer) *svg.SVG {
	canvas := svg.New(writer)
	canvas.Start(width, height)

	// positionsA := newPositions(g)

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

			nodePos := &Vector{
				X:   int((width / 2) + math.Cos(angle)*rad),
				Y:   int((height / 2) + math.Sin(angle)*rad),
				Rot: (angle + rotationCorrection) * (180 / math.Pi),
				Rad: rad,
			}

			canvas.Circle(
				nodePos.X,
				nodePos.Y,
				1,
				"fill:black")

			// for neighbor, _ := range g.Nodes[node] {
			// 	neighborPos := positionsA.SafeGet(neighbor)

			// 	canvas.Line(
			// 		nodePos.X, nodePos.Y,
			// 		neighborPos.X, neighborPos.Y,
			// 		"stroke:rgba(0, 0, 0, 0.2);stroke-width:0.5")
			// }

			canvas.Text(
				0,
				0,
				node.Name,
				`style="text-anchor:middle;font-size:12px;fill:#5bb4c0;"`,
				fmt.Sprintf(`transform="rotate(%v, %v, %v) translate(%v, %v)"`, nodePos.Rot, nodePos.X, nodePos.Y, nodePos.X, nodePos.Y))

			angle += (math.Pi * 2) / float64(len(nodes))
		}
	}

	canvas.End()

	return canvas
}
