package main

import (
	"log"
	"os"

	"github.com/aaasen/kapok"
)

func main() {
	in, err := os.Open("/home/aasen/dev/data/enwiki-latest-pages-articles.xml")

	if err != nil {
		log.Fatal(err)
	}

	kapok.GenerateNeoGraph(in, -1)
}
