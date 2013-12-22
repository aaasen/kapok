package parse

import (
	"bytes"
	"errors"
)

// Page is a representation of a Wikipedia page with only the necessary fields.
// A Wikipedia page can be unmarshalled into a page just fine.
type Page struct {
	Title      string    `xml:"title"`
	Revision   *Revision `xml:"revision"`
	Links      []string
	Categories []string
}

func (self *Page) String() string {
	return self.Title
}

var (
	ErrTitleNotFound = errors.New("error parsing page from xml: title not found")
)

func NewPageFromXML(text []byte) (*Page, error) {
	page := &Page{}

	err := page.getTitle(text)

	if err != nil {
		return nil, err
	}

	page.getLinks(text)

	return page, nil
}

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
				page.Links = append(page.Links, string(linkBody))
			}

			text = text[endIndex+len(endTag):]
		} else {
			linksLeft = false
		}
	}
}

func (page *Page) getTitle(text []byte) error {
	startTag := []byte("<title>")
	startIndex := bytes.Index(text, startTag)

	endTag := []byte("</title>")
	endIndex := bytes.Index(text, endTag)

	if startIndex != -1 && endIndex != -1 {
		title := text[startIndex+len(startTag) : endIndex]

		if len(title) < 1 {
			return ErrTitleNotFound
		}

		page.Title = string(title)
	} else {
		return ErrTitleNotFound
	}

	return nil
}

func isTitle(title []byte) bool {
	specialIndex := bytes.IndexAny(title, ":#{}/")

	return specialIndex == -1
}

// removeEscapedRegions removes all markup in between the <nowiki> tags.
// See http://www.mediawiki.org/wiki/Help:Formatting
func removeEscapedRegions(page []byte) []byte {
	return page
}

// Revision usually contains information about the user and time of the revision.
// Since Kapok is focused only on the latest version of Wikipedia, these fields are ignored.
type Revision struct {
	Text string `xml:"text"`
}
