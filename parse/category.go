package parse

import (
	"fmt"
	"strings"
)

type Categories struct {
	categories map[string][]*Page
}

func NewCategories() *Categories {
	return &Categories{
		categories: make(map[string][]*Page),
	}
}

func (self *Categories) AddPage(page *Page, cats []string) {
	for _, cat := range cats {
		self.categories[cat] = append(self.categories[cat], page)
	}
}

func (self *Categories) String() string {
	str := ""

	for category, pages := range self.categories {
		str += fmt.Sprintf("%s -> (", category)

		pageStrings := make([]string, len(pages))
		for i, page := range pages {
			pageStrings[i] = page.String()
		}

		str += strings.Join(pageStrings, ", ")
		str += ")\n"
	}

	return str
}
