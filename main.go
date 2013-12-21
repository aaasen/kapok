package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}

	switch os.Args[1] {
	case "generate":
		// "/home/aasen/dev/data/enwiki-latest-pages-articles.xml"
		// "/home/aasen/downloads/articles.csv"
		// "/home/aasen/downloads/rels.csv"

		if len(os.Args) != 5 {
			fmt.Println("generate requires exactly 3 arguments: xml path, nodes path, rels path")
			return
		}

		err := GenerateByPath(os.Args[2], os.Args[3], os.Args[4])

		if err != nil {
			fmt.Println(err)
		}
	}
}
