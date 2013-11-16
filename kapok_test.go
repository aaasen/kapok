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
	numPages := 0
	maxPages := 10

	for {
		select {
		case page := <-pages:
			log.Println(numPages)

			origin := graph.SafeGet(page.Title)

			for _, dest := range page.Links {
				graph.AddArc(origin, graph.SafeGet(dest))
			}

			numPages++

			if numPages >= maxPages {
				file, err := os.Create("/home/aasen/dev/data/wiki-graph.gob")

				if err != nil {
					log.Fatal("error opening graph file: ", err)
				}

				err = graph.Export(file)

				if err != nil {
					log.Fatal("error exporting graph: ", err)
				}

				return
			}
		}
	}
}
