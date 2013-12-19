package kapok

import (
	"io"
	"log"
	"time"

	"github.com/aaasen/kapok/parse"

	"github.com/jmcvetta/neoism"
)

type governor struct {
	maxConnections     int
	currentConnections int
	queryChan          chan *neoism.CypherQuery
	ticker             <-chan time.Time
	db                 *neoism.Database
}

func (self *governor) run() {
	for _ = range self.ticker {
		if self.currentConnections < self.maxConnections {
			select {
			case query, ok := <-self.queryChan:
				self.currentConnections++

				if !ok {
					return
				}

				err := self.db.Cypher(query)

				if err != nil {
					log.Println("error executing cypher query: " + err.Error())
				}

				self.currentConnections--
			}
		}
	}
}

// GenerateNeoGraph transfers a Wikipedia database dump into a Neo4J database.
// in is a reader that contains the database,
// maxPages is the maxiumum number of pages to load (or -1 for unlimited)
func GenerateNeoGraph(in io.Reader, maxPages int) error {
	graph, err := neoism.Connect("http://localhost:7474/db/data")

	if err != nil {
		return err
	}

	governorTicker := time.Tick(time.Millisecond * 10)
	queryChan := make(chan *neoism.CypherQuery, 512)

	governor := &governor{
		maxConnections:     512,
		currentConnections: 0,
		queryChan:          queryChan,
		ticker:             governorTicker,
		db:                 graph,
	}

	pages := make(chan *parse.Page, 1024)
	go parse.CategorizedParse(in, pages)

	go governor.run()

	start := time.Now()
	numPages := 0

	for {
		select {
		case page, ok := <-pages:
			if !ok {
				return nil
			}

			if maxPages != -1 && numPages >= maxPages {
				return nil
			}

			if numPages%1 == 0 {
				log.Printf("processed %v pages in %v", numPages, time.Since(start))
			}

			numPages++

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

			queryChan <- &cypherQuery
		}
	}
}
