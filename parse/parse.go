package parse

// An ad-hoc parser for Wikipedia's 40GB (and growing) XML database.

import (
	"bufio"
	"bytes"
	"io"
	"log"
)

type Parser struct {
	BytesProcessed int64
}

func NewParser() *Parser {
	return &Parser{
		0,
	}
}

// Parse parses given reader as XML and dumps Page objects into the given channel.
// Parse will fill the Page's Title, Links, and Categories.
//
// When the reader is empty, Parse will close its output channel.
//
// Parse will throw away malformed input instead of exiting and reporting it.
func (parser *Parser) Parse(reader io.Reader, pages chan<- *Page) {
	rawPages := make(chan []byte)

	go parser.getRawPages(reader, rawPages)
	go parser.getPages(rawPages, pages)
}

// getRawPages creates full pages from a reader that can then be parsed with an XML parser.
func (parser *Parser) getRawPages(rawReader io.Reader, pages chan<- []byte) {
	reader := bufio.NewReader(rawReader)

	buffer := make([]byte, 0)
	inPage := false

	eof := false

	for !eof {
		text, err := reader.ReadBytes('>')
		parser.BytesProcessed += int64(len(text))

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

			if bytes.Index(buffer, []byte("#REDIRECT")) == -1 {
				pages <- buffer
			}

			buffer = make([]byte, 0)
		} else if inPage {
			buffer = append(buffer, text...)
		}
	}

	close(pages)
}

// GetPages parses a complete XML page into a page object.
func (parser *Parser) getPages(rawPages <-chan []byte, pages chan<- *Page) {
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
