package generate

import (
	"os"
	"testing"

	"github.com/aaasen/kapok/parse"
)

func BenchmarkGenerate(b *testing.B) {
	nodes, err := os.Create(".kapok-benchmark-nodes.csv")

	if err != nil {
		b.Error(err)
	}

	rels, err := os.Create(".kapok-benchmark-rels.csv")

	if err != nil {
		b.Error(err)
	}

	page := &parse.Page{
		Title:      "origin",
		Links:      []string{"a", "b", "c"},
		Categories: []string{"cat0", "cat1", "cat2"},
	}

	gen := generate.NewCSVGenerator()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		gen.GeneratePage(page, nodes, rels)
	}
}
