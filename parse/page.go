package parse

import (
	"bytes"
	"errors"
)

// Page is a representation of a Wikipedia page with only the necessary fields.
type Page struct {
	Title      string
	Links      []string
	Categories []string
}

var (
	ErrTitleNotFound = errors.New("error parsing page from xml: title not found")
)

func (self *Page) String() string {
	return self.Title
}

// NewPageFromXML creates a Page object from XML.
func NewPageFromXML(text []byte) (*Page, error) {
	page := &Page{}

	err := page.getTitle(text)

	if err != nil {
		return nil, err
	}

	page.getLinks(text)

	return page, nil
}

// getLinks populates a Page's Links and Categories fields by parsing the given XML.
// It will only take into account internal links, piped links, and categories.
// See http://www.mediawiki.org/wiki/Help:Links for the syntax of these links.
func (page *Page) getLinks(text []byte) {
	linksLeft := true

	for linksLeft {
		startTag := []byte("[[")
		startIndex := bytes.Index(text, startTag)

		endTag := []byte("]]")
		endIndex := bytes.Index(text, endTag)

		if startIndex != -1 && endIndex != -1 && startIndex+len(startTag) < endIndex {
			linkBody := text[startIndex+len(startTag) : endIndex]

			pipeIndex := bytes.Index(linkBody, []byte("|"))

			if pipeIndex != -1 {
				linkBody = linkBody[:pipeIndex]
			}

			categoryTag := []byte("Category:")
			categoryIndex := bytes.Index(linkBody, categoryTag)

			if categoryIndex != -1 {
				category := linkBody[len(categoryTag):]

				page.Categories = append(page.Categories, string(category))
			} else {
				if isTitle(linkBody) {
					page.Links = append(page.Links, string(linkBody))
				}
			}

			text = text[endIndex+len(endTag):]
		} else {
			linksLeft = false
		}
	}
}

// getTitle parses the title from an XML representation of a Wikipedia page.
// In the event of an error or malformed XML, it will return ErrTitleNotFound.
func (page *Page) getTitle(text []byte) error {
	startTag := []byte("<title>")
	startIndex := bytes.Index(text, startTag)

	endTag := []byte("</title>")
	endIndex := bytes.Index(text, endTag)

	if startIndex != -1 && endIndex != -1 {
		title := text[startIndex+len(startTag) : endIndex]

		if len(title) < 1 || !isTitle(title) {
			return ErrTitleNotFound
		}

		page.Title = string(title)
	} else {
		return ErrTitleNotFound
	}

	return nil
}

// isTitle returns whether or not the given byte array looks like a valid
// title for a Wikipedia article.s
// It is based off of http://www.mediawiki.org/wiki/Help:Links
// and tries to only accept internal links.
func isTitle(title []byte) bool {
	specialIndex := bytes.IndexAny(title, ":#{}/")

	return specialIndex == -1
}

// removeEscapedRegions removes all markup in between the <nowiki> tags.
// See http://www.mediawiki.org/wiki/Help:Formatting
func removeEscapedRegions(page []byte) []byte {
	return page
}
