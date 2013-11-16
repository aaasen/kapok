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
var linkRegex = regexp.MustCompile("\\[\\[([^|]+?)\\]\\]")
var categoryRegex = regexp.MustCompile("\\[\\[Category:(.+?)\\]\\]")

func Parse(reader io.Reader, pages chan<- *Page) {
	chunks := make(chan []byte)
	rawPages := make(chan []byte)
	somePages := make(chan *Page)

	go getChunks(reader, chunks)
	go getRawPages(chunks, rawPages)
	go getPages(rawPages, somePages)
	go getLinks(somePages, pages)

}

func CategorizedParse(reader io.Reader, out chan<- *Page, categories *Categories) {
	pages := make(chan *Page)

	go getCategories(pages, out, categories)

	Parse(reader, pages)
}

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

func getCategories(pages <-chan *Page, categorizedPages chan<- *Page, categories *Categories) {
	for {
		select {
		case page := <-pages:
			rawCats := categoryRegex.FindAllStringSubmatch(page.Revision.Text, -1)
			cats := make([]string, len(rawCats))

			for i, rawCat := range rawCats {
				cats[i] = rawCat[1]
			}

			categories.AddPage(page, cats)

			categorizedPages <- page
		}
	}
}
