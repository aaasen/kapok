package kapok

import (
	"log"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/aaasen/kapok/graph"
)

func TestGenerateGraph(t *testing.T) {
	Convey("Generating a graph should work", t, func() {
		in, err := os.Open("/home/aasen/dev/data/enwiki-latest-pages-articles.xml")

		if err != nil {
			log.Fatal(err)
		}

		var g *graph.Graph

		So(func() {
			g = GenerateGraph(in, 100)
		}, ShouldNotPanic)

		So(len(g.Nodes), ShouldBeGreaterThanOrEqualTo, 100)
	})
}

func TestGenerateAndStore(t *testing.T) {
	Convey("Generating and storing should work", t, func() {
		So(func() {
			in, err := os.Open("/home/aasen/dev/data/enwiki-latest-pages-articles.xml")

			if err != nil {
				t.Fatal("error opening database file: ", err)
			}

			err = GenerateAndStore(in, "/home/aasen/dev/data/wiki-graph.gob", 10, 5)

			if err != nil {
				t.Fatal("error in GenerateAndStore: ", err)
			}
		}, ShouldNotPanic)
	})
}
