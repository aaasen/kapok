package parse

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
