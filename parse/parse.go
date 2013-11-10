package parse

import (
	"bufio"
	"io"
	"log"
	"regexp"
)

var pageStartRegex = regexp.MustCompile(".*<page>.*")
var pageEndRegex = regexp.MustCompile(".*</page>.*")

func getChunks(reader io.Reader, chunks chan<- []byte) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		chunks <- scanner.Bytes()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getPages(chunks <-chan []byte, pages chan<- []byte) {
	page := make([]byte, 0)
	inPage := false

	for {
		select {
		case chunk := <-chunks:
			if pageStartRegex.Match(chunk) {
				inPage = true
			} else if pageEndRegex.Match(chunk) {
				if inPage {
					pages <- page
					page = make([]byte, 0)
				}

				inPage = false
			} else {
				if inPage {
					page = append(page, chunk...)
				}
			}
		}
	}
}

func getXML(pages <-chan []byte) {
	for {
		select {
		case page := <-pages:
			log.Println(string(page))
		}
	}
}
