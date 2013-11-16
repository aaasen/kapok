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

	numPages := 0
	maxPages := 10

	for {
		select {
		case <-pages:
			numPages++

			log.Println(numPages)

			if numPages >= maxPages {
				return
			}
		}
	}
}

func TestCategorizedParse(t *testing.T) {
	file, err := os.Open("/home/aasen/dev/data/enwiki-latest-pages-articles.xml")

	if err != nil {
		t.Error(err)
	}

	pages := make(chan *Page)
	cats := NewCategories()
	defer log.Println(cats)

	go CategorizedParse(file, pages, cats)

	numPages := 0
	maxPages := 10

	for {
		select {
		case <-pages:
			numPages++

			if numPages >= maxPages {
				return
			}
		}
	}
}
