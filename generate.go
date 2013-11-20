package kapok

import (
	"io"

	"github.com/aaasen/kapok/graph"
	"github.com/aaasen/kapok/parse"
)

// GenerateGraph reads the contents of the given reader and creates
// a graph from the first maxPages pages.
//
// If maxPages is -1, GenerateGraph will read until the database channel closes.
func GenerateGraph(in io.Reader, maxPages int) *graph.Graph {
	pages := make(chan *parse.Page)
	go parse.Parse(in, pages)

	graph := graph.NewGraph()
	numPages := 0

	for {
		select {
		case page, ok := <-pages:
			if !ok {
				return graph
			}

			origin := graph.SafeGet(page.Title)

			for _, dest := range page.Links {
				graph.AddArc(origin, graph.SafeGet(dest))
			}

			if numPages > maxPages {
				return graph
			}

			if maxPages != -1 {
				numPages++
			}
		}
	}
}
