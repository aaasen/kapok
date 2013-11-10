package parse

type Page struct {
	Title    string    `xml:"title"`
	Revision *Revision `xml:"revision"`
	Links    []string
}

type Revision struct {
	Text string `xml:"text"`
}
