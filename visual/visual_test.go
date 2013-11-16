package visual

import (
	"log"
	"os"
	"testing"

	"github.com/aaasen/kapok/graph"
)

func TestVisual(t *testing.T) {
	in, err := os.Open("/home/aasen/dev/data/wiki-graph.gob")

	if err != nil {
		t.Fatal("error opening graph file: ", err)
	}

	out, err := os.Create("/home/aasen/dev/data/graph.svg")

	if err != nil {
		t.Fatal("error creating svg file: ", out)
	}

	err, g := graph.Import(in)

	if err != nil {
		log.Fatal("error importing graph: ", err)
	}

	Visualise(g, out)
}
