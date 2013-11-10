package parse

import (
	"os"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	file, err := os.Open("/home/aasen/dev/data/enwiki-latest-pages-articles.xml")

	if err != nil {
		t.Error(err)
	}

	chunks := make(chan []byte)
	rawPages := make(chan []byte)
	pages := make(chan *Page)
	linkedPages := make(chan *Page)

	go getChunks(file, chunks)
	go getRawPages(chunks, rawPages)
	go getPages(rawPages, pages)
	go getLinks(pages, linkedPages)
	go printPages(linkedPages)

	time.Sleep(time.Minute)

}
