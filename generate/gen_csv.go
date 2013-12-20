package generate

import (
	"fmt"
	"io"

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
		title = "cat:" + title
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
		nodes.Write([]byte(fmt.Sprintf("%s\t%s\n", page.Title, "Article")))
	}

	for _, link := range page.Links {
		linkId, created := self.ids.GetID(link, false)

		if created {
			nodes.Write([]byte(fmt.Sprintf("%s\t%s\n", link, "Article")))
		}

		rels.Write([]byte(fmt.Sprintf("%d\t%d\t%s\n", originId, linkId, "REFERS_TO")))
	}

	for _, category := range page.Categories {
		catId, created := self.ids.GetID(category, true)

		if created {
			nodes.Write([]byte(fmt.Sprintf("%s\t%s\n", category, "Category")))
		}

		rels.Write([]byte(fmt.Sprintf("%d\t%d\t%s\n", originId, catId, "IN_CATEGORY")))
	}
}
