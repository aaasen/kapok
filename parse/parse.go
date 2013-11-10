package parse

import (
	"bufio"
	"encoding/xml"
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

func getRawPages(chunks <-chan []byte, pages chan<- []byte) {
	page := []byte("<page>")
	inPage := false

	for {
		select {
		case chunk := <-chunks:
			if pageStartRegex.Match(chunk) {
				inPage = true
			} else if pageEndRegex.Match(chunk) {
				if inPage {
					page = append(page, []byte("</page>")...)
					pages <- page
					page = []byte("<page>")
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

func getPages(rawPages <-chan []byte, pages chan<- *page) {
	for {
		select {
		case rawPage := <-rawPages:
			pageStruct := &page{}

			err := xml.Unmarshal(rawPage, pageStruct)

			if err != nil {
				log.Println(string(rawPage))
				log.Println(err)
			} else {
				pages <- pageStruct
				log.Println(pageStruct)
			}
		}
	}
}

func printPages(pages <-chan *page) {
	for {
		select {
		case page := <-pages:
			log.Println(page)
		}
	}
}
