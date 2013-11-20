package kapok

import (
	"io"
	"os"

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

// GenerateAndStore generates a graph from the given reader and
// continually exports it to disk.
//
// It will only consume maxPages pages, or unlimited if maxPages is -1
//
// saveInterval represents the number of pages between each export.
func GenerateAndStore(in io.Reader, outPath string, maxPages, saveInterval int) error {
	pages := make(chan *parse.Page)
	go parse.Parse(in, pages)

	graph := graph.NewGraph()
	numPages := 0

	for {
		select {
		case page, ok := <-pages:
			numPages++

			if !ok {
				err := storeGraph(graph, outPath)

				if err != nil {
					return err
				}
			}

			origin := graph.SafeGet(page.Title)

			for _, dest := range page.Links {
				graph.AddArc(origin, graph.SafeGet(dest))
			}

			if maxPages != -1 && numPages >= maxPages {
				err := storeGraph(graph, outPath)

				return err
			}

			if numPages%saveInterval == 0 {
				err := storeGraph(graph, outPath)

				if err != nil {
					return err
				}
			}
		}
	}
}

func storeGraph(g *graph.Graph, outPath string) error {
	err := os.Remove(outPath)

	if err != nil {
		return err
	}

	out, err := os.Create(outPath)

	if err != nil {
		return err
	}

	g.Export(out)

	return nil
}
