package parse

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	file, err := os.Open("/home/aasen/dev/data/enwiki-latest-pages-articles.xml")

	if err != nil {
		t.Error(err)
	}

	pages := make(chan []byte)

	chunk(file, pages)
}
