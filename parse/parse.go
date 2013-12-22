// An ad-hoc parser for Wikipedia's 45GB (and growing) XML database.
package parse

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"regexp"
)

var pageStartRegex = regexp.MustCompile(".*<page>.*")
var pageEndRegex = regexp.MustCompile(".*</page>.*")
var redirectRegex = regexp.MustCompile("#REDIRECT[ \t].*?\\[\\[.*?\\]\\]")
var linkRegex = regexp.MustCompile("\\[\\[([^|]+?)\\]\\]")
var categoryRegex = regexp.MustCompile("\\[\\[Category:(.+?)\\]\\]")

// Parse parses given reader as XML and dumps Page objects with links
// into its output channel.
func Parse(reader io.Reader, pages chan<- *Page) {
	rawPages := make(chan []byte)
	nonRedirectPages := make(chan []byte)

	go GetRawPages(reader, rawPages)
	go FilterRedirects(rawPages, nonRedirectPages)
	go GetPages(nonRedirectPages, pages)
}

// GetRawPages creates full pages from a reader that can then be parsed with an XML parser.
func GetRawPages(rawReader io.Reader, pages chan<- []byte) {
	reader := bufio.NewReader(rawReader)

	buffer := make([]byte, 0)
	inPage := false

	eof := false

	for !eof {
		text, err := reader.ReadBytes('>')

		if err != nil {
			if err == io.EOF {
				eof = true
			} else {
				log.Println(err.Error() + " skipping line")
			}
		}

		startTag := []byte("<page>")
		startIndex := bytes.Index(text, startTag)

		endTag := []byte("</page>")
		endIndex := bytes.Index(text, endTag)

		if startIndex != -1 {
			inPage = true
			buffer = text[startIndex:]
		} else if endIndex != -1 {
			inPage = false
			buffer = append(buffer, text[:endIndex+len(endTag)]...)

			pages <- buffer

			buffer = make([]byte, 0)
		} else if inPage {
			buffer = append(buffer, text...)
		}
	}

	close(pages)
}

// FilterRedirects discards all pages that redirect to another page.
func FilterRedirects(rawPages <-chan []byte, nonRedirectPages chan<- []byte) {
	for {
		select {
		case rawPage, ok := <-rawPages:
			if !ok {
				close(nonRedirectPages)
				return
			}

			if redirectRegex.Find(rawPage) == nil {
				nonRedirectPages <- rawPage
			}
		}
	}
}

// GetPages parses a complete XML page into a page object.
func GetPages(rawPages <-chan []byte, pages chan<- *Page) {
	for {
		select {
		case rawPage, ok := <-rawPages:
			if !ok {
				close(pages)
				return
			}

			page, err := NewPageFromXML(rawPage)

			if err == nil {
				pages <- page
			}
		}
	}
}
