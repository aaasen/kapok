package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/aaasen/kapok/generate"
	"github.com/aaasen/kapok/parse"
)

func GenerateByPath(inPath, nodesPath, relsPath string) error {
	in, err := os.Open(inPath)

	if err != nil {
		return err
	}

	nodes, err := os.Create(nodesPath)

	if err != nil {
		return err
	}

	rels, err := os.Create(relsPath)

	if err != nil {
		return err
	}

	Generate(in, nodes, rels)

	return nil
}

func Generate(in io.Reader, nodes io.Writer, rels io.Writer) {

	nodes.Write([]byte("i:id\ttitle\tl:label\n"))
	rels.Write([]byte("start\tend\ttype\n"))

	gen := generate.NewCSVGenerator()

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

			gen.GeneratePage(page, nodes, rels)

			if numPages%100 == 0 {
				log.Printf("processed %v pages in %v (%.2f pages/sec)",
					numPages, time.Since(start), float64(numPages)/time.Since(start).Seconds())
			}

			numPages++
		}
	}
}
