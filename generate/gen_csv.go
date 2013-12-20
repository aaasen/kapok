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
		1,
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
	articles io.Writer, categories io.Writer, rels io.Writer) {

	originId, created := self.ids.GetID(page.Title, false)

	if created {
		articles.Write([]byte(page.Title + "\n"))
	}

	for _, link := range page.Links {
		linkId, created := self.ids.GetID(link, false)

		if created {
			articles.Write([]byte(link + "\n"))
		}

		rels.Write([]byte(fmt.Sprintf("%d\t%d\t%s\n", originId, linkId, "REFERS_TO")))
	}

	for _, category := range page.Categories {
		catId, created := self.ids.GetID(category, true)

		if created {
			categories.Write([]byte(category + "\n"))
		}

		rels.Write([]byte(fmt.Sprintf("%d\t%d\t%s\n", originId, catId, "IN_CATEGORY")))
	}
}
