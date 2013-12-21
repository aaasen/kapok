package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/aaasen/kapok/generate"
)

var cpuprofile = true
var memprofile = true

func main() {
	if len(os.Args) <= 1 {
		return
	}

	if cpuprofile {
		f, err := os.Create("kapok_cpu.prof")

		if err != nil {
			log.Fatal(err)
		}

		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	switch os.Args[1] {
	case "generate":
		if len(os.Args) < 5 {
			log.Fatal("generate requires at least 3 arguments: xml path, nodes path, rels path")
			return
		}

		err := generateByPath(os.Args[2], os.Args[3], os.Args[4], 10000)

		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Println("command not found")
	}

	if memprofile {
		f, err := os.Create("kapok_mem.prof")

		if err != nil {
			log.Fatal(err)
		}

		pprof.WriteHeapProfile(f)
		f.Close()
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
