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
	pages := make(chan []byte)

	go getChunks(file, chunks)
	go getPages(chunks, pages)
	go getXML(pages)

	time.Sleep(time.Minute)

}
