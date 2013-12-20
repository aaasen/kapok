package generate

import (
	"os"
	"testing"

	"github.com/aaasen/kapok/parse"
)

func TestGeneratePage(t *testing.T) {
	page := &parse.Page{
		Title:      "origin",
		Links:      []string{"a", "b", "c"},
		Categories: []string{"cat0", "cat1", "cat2"},
	}

	nodes, _ := os.Create("articles.csv")
	rels, _ := os.Create("rels.csv")

	gen := NewCSVGenerator()

	gen.GeneratePage(page, nodes, rels)
}
