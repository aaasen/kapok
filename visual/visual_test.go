package visual

import (
	"os"
	"testing"

	"github.com/aaasen/kapok/graph"
)

func TestVisual(t *testing.T) {
	in, err := os.Open("/home/aasen/dev/data/wiki-graph.gob")

	if err != nil {
		t.Fatal("error importing graph: ", err)
	}

	out, err := os.Create("/home/aasen/dev/data/graph.svg")

	if err != nil {
		t.Fatal("error creating svg file: ", out)
	}

	g := graph.Import(in)

	Visualise(g, out)
}
