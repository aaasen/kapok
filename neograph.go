package kapok

import (
	"errors"
	"io"
	"log"
	"time"

	"github.com/aaasen/kapok/parse"

	"github.com/jmcvetta/neoism"
)

func GenerateNeoGraph(in io.Reader, maxPages int) error {
	graph, err := neoism.Connect("http://localhost:7474/db/data")

	if err != nil {
		return err
	}

	pages := make(chan *parse.Page)
	go parse.Parse(in, pages)

	start := time.Now()
	numPages := 0

	for {
		select {
		case page, ok := <-pages:
			numPages++

			if !ok {
				return nil
			}

			origin, _, err := graph.GetOrCreateNode("pages", "title", neoism.Props{"title": page.Title})

			if err != nil {
				return errors.New("error fetching origin: " + err.Error())
			}

			for _, destTitle := range page.Links {

				dest, _, err := graph.GetOrCreateNode("pages", "title", neoism.Props{"title": destTitle})

				if err != nil {
					return errors.New("error fetching dest: " + err.Error())
				}

				_, err = origin.Relate("link", dest.Id(), neoism.Props{})

				if err != nil {
					return errors.New("error relating origin and dest: " + err.Error())
				}
			}

			if maxPages != -1 && numPages >= maxPages {
				return nil
			}

			if numPages%100 == 0 {
				log.Println("processed %v page in %v", numPages, time.Since(start))
			}
		}
	}
}
