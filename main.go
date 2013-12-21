package main

import (
	"fmt"
	"os"

	"github.com/aaasen/kapok/generate"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}

	switch os.Args[1] {
	case "generate":
		if len(os.Args) != 5 {
			fmt.Println("generate requires exactly 3 arguments: xml path, nodes path, rels path")
			return
		}

		err := generateByPath(os.Args[2], os.Args[3], os.Args[4], -1)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func generateByPath(inPath, nodesPath, relsPath string, maxPages int) error {
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

	generate.Generate(in, nodes, rels, maxPages)

	return nil
}
