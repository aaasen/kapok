package kapok

import (
	"log"
	"os"
	"testing"

	"github.com/aaasen/kapok/graph"
	"github.com/aaasen/kapok/parse"
)

func TestKapok(t *testing.T) {
	file, err := os.Open("/home/aasen/dev/data/enwiki-latest-pages-articles.xml")

	if err != nil {
		log.Fatal(err)
	}

	pages := make(chan *parse.Page)
	go parse.Parse(file, pages)

	graph := graph.NewGraph()

	for {
		select {
		case page := <-pages:
			origin := graph.SafeGet(page.Title)

			for _, dest := range page.Links {
				graph.AddArc(origin, graph.SafeGet(dest))
			}

			log.Println(graph.String())
		}
	}
}
