package kapok

import (
	"io"
	"log"
	"time"

	"github.com/aaasen/kapok/parse"

	"github.com/jmcvetta/neoism"
)

// GenerateNeoGraph transfers a Wikipedia database dump into a Neo4J database.
// in is a reader that contains the database,
// maxPages is the maxiumum number of pages to load (or -1 for unlimited)
func GenerateNeoGraph(in io.Reader, maxPages int) error {
	graph, err := neoism.Connect("http://localhost:7474/db/data")

	if err != nil {
		return err
	}

	pages := make(chan *parse.Page)
	go parse.CategorizedParse(in, pages)

	start := time.Now()
	numPages := 0

	for {
		select {
		case page, ok := <-pages:
			if maxPages != -1 && numPages >= maxPages {
				return nil
			}

			if numPages%1 == 0 {
				log.Printf("processed %v pages in %v", numPages, time.Since(start))
			}

			numPages++

			if !ok {
				return nil
			}

			result := []struct {
				N neoism.Node
			}{}

			cypherQuery := neoism.CypherQuery{
				Statement: `
				MERGE (origin:Article { title:{pageTitle} })
				FOREACH (otherName IN {otherNames} |
					MERGE (other:Article { title: otherName})
					CREATE UNIQUE (origin)-[r:REFERS_TO]->(other))
				FOREACH (catName IN {catNames} |
					MERGE (cat:Category { title:catName }) 
					CREATE UNIQUE (origin)-[:IN_CATEGORY]-(cat)) 
				RETURN origin
				`,
				Parameters: neoism.Props{
					"pageTitle":  page.Title,
					"catNames":   page.Categories,
					"otherNames": page.Links},
				Result: &result,
			}

			go func() {
				err := graph.Cypher(&cypherQuery)

				if err != nil {
					log.Println("error executing cypher query: " + err.Error())
				}
			}()
		}
	}
}
