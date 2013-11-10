package parse

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"regexp"
)

var pageStartRegex = regexp.MustCompile(".*<page>.*")
var pageEndRegex = regexp.MustCompile(".*</page>.*")
var linkRegex = regexp.MustCompile("\\[\\[([^|]+?)\\]\\]")

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

func getPages(rawPages <-chan []byte, pages chan<- *Page) {
	for {
		select {
		case rawPage := <-rawPages:
			pageStruct := &Page{}

			err := xml.Unmarshal(rawPage, pageStruct)

			if err != nil {
				log.Println(string(rawPage))
				log.Println(err)
			} else {
				pages <- pageStruct
			}
		}
	}
}

func getLinks(pages <-chan *Page, linkedPages chan<- *Page) {
	for {
		select {
		case page := <-pages:
			links := linkRegex.FindAllStringSubmatch(page.Revision.Text, -1)

			for _, link := range links {
				page.Links = append(page.Links, link[1])
			}

			linkedPages <- page
		}
	}
}

func printPages(pages <-chan *Page) {
	for {
		select {
		case page := <-pages:
			fmt.Sprint(page)
		}
	}
}
