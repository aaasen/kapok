package parse

type Page struct {
	Title    string    `xml:"title"`
	Revision *Revision `xml:"revision"`
	Links    []string
}

func (self *Page) String() string {
	return self.Title
}

type Revision struct {
	Text string `xml:"text"`
}
