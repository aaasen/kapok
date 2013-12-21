package generate

import (
	"fmt"
	"io"
	"regexp"

	"github.com/aaasen/kapok/parse"
)

type CSVGenerator struct {
	ids *IDGenerator
}

func NewCSVGenerator() *CSVGenerator {
	return &CSVGenerator{
		NewIDGenerator(),
	}
}

type IDGenerator struct {
	currentId int64
	labels    map[string]int64
}

func NewIDGenerator() *IDGenerator {
	return &IDGenerator{
		0,
		make(map[string]int64),
	}
}

func (self *IDGenerator) GetID(title string, category bool) (int64, bool) {
	if category {
		title = "c4ef35:" + title
	}

	id, exists := self.labels[title]

	if exists {
		return id, false
	}

	self.labels[title] = self.currentId
	self.currentId++

	return self.currentId - 1, true
}

func (self *CSVGenerator) GeneratePage(page *parse.Page,
	nodes io.Writer, rels io.Writer) {

	originId, created := self.ids.GetID(page.Title, false)

	if created {
		writeNode(nodes, originId, page.Title, "Article")
	}

	for _, link := range page.Links {
		linkId, created := self.ids.GetID(link, false)

		if created {
			writeNode(nodes, linkId, link, "Article")
		}

		writeRel(rels, originId, linkId, "REFERS_TO")
	}

	for _, category := range page.Categories {
		catId, created := self.ids.GetID(category, true)

		if created {
			writeNode(nodes, catId, category, "Category")
		}

		writeRel(rels, originId, catId, "IN_CATEGORY")
	}
}

var invalidCharacterRegex = regexp.MustCompile("[\t\"\\']")

func writeNode(out io.Writer, id int64, title string, label string) {
	title = invalidCharacterRegex.ReplaceAllString(title, "")

	out.Write([]byte(fmt.Sprintf("%d\t%s\t%s\n", id, title, label)))
}

func writeRel(out io.Writer, origin int64, dest int64, rel string) {
	out.Write([]byte(fmt.Sprintf("%d\t%d\t%s\n", origin, dest, rel)))
}
