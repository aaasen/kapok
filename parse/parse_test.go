package parse

import (
	"log"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	file, err := os.Open("/home/aasen/dev/data/enwiki-latest-pages-articles.xml")

	if err != nil {
		t.Error(err)
	}

	pages := make(chan *Page)

	go Parse(file, pages)

	i := 1

	for {
		select {
		case <-pages:
			log.Println(i)
			i++
		}
	}
}
