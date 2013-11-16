package parse

// Page is a representation of a Wikipedia page with only the necessary fields.
// A Wikipedia page can be unmarshalled into a page just fine.
type Page struct {
	Title    string    `xml:"title"`
	Revision *Revision `xml:"revision"`
	Links    []string
}

func (self *Page) String() string {
	return self.Title
}

// Revision usually contains information about the user and time of the revision.
// Since Kapok is focused only on the latest version of Wikipedia, these fields are ignored.
type Revision struct {
	Text string `xml:"text"`
}
