package parse

import (
	"fmt"
	"strings"
)

// Categories is a simple structure for keeping track of the categories of Wikipedia articles.
// It is optimized for queries like "articles about science"
// rather than "which categories is this article in".
type Categories struct {
	categories map[string][]*Page
}

// NewCategories returns an empty categories object.
func NewCategories() *Categories {
	return &Categories{
		categories: make(map[string][]*Page),
	}
}

// AddPage adds the given page to the given categories.
func (self *Categories) AddPage(page *Page, cats []string) {
	for _, cat := range cats {
		self.categories[cat] = append(self.categories[cat], page)
	}
}

// String produces a string represenation of the categories in the form:
// category -> (article, article, article)
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
