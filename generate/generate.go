package generate

import (
	"io"
	"log"
	"time"

	"github.com/aaasen/kapok/parse"
)

func Generate(in io.Reader, nodes io.Writer, rels io.Writer, maxPages int) {

	nodes.Write([]byte("i:id\ttitle\tl:label\n"))
	rels.Write([]byte("start\tend\ttype\n"))

	gen := NewCSVGenerator()

	pages := make(chan *parse.Page)
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

			if numPages > maxPages && maxPages != -1 {
				return
			}

			gen.GeneratePage(page, nodes, rels)

			if numPages%100 == 0 {
				log.Printf("processed %v pages in %v (%.2f pages/sec)",
					numPages, time.Since(start), float64(numPages)/time.Since(start).Seconds())
			}

			page = nil
			numPages++
		}
	}
}
