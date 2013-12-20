package main

import (
	"log"
	"os"
	"time"

	"github.com/aaasen/kapok/generate"
	"github.com/aaasen/kapok/parse"
)

func main() {
	in, err := os.Open("/home/aasen/dev/data/enwiki-latest-pages-articles.xml")

	if err != nil {
		log.Fatal(err)
	}

	nodes, _ := os.Create("/home/aasen/downloads/articles.csv")
	nodes.Write([]byte("title\tl:label\n"))

	rels, _ := os.Create("/home/aasen/downloads/rels.csv")
	rels.Write([]byte("start\tend\ttype\n"))

	gen := generate.NewCSVGenerator()

	pages := make(chan *parse.Page, 1024)
	go parse.CategorizedParse(in, pages)

	start := time.Now()
	numPages := 0

	for {
		select {
		case page, ok := <-pages:
			if !ok {
				log.Println("channel close, exiting")
				return
			}

			gen.GeneratePage(page, nodes, rels)

			if numPages%100 == 0 {
				log.Printf("processed %v pages in %v", numPages, time.Since(start))
			}

			numPages++
		}
	}
}
