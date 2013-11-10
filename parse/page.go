package parse

type page struct {
	Title    string    `xml:"title"`
	Revision *revision `xml:"revision"`
}

type revision struct {
	Text string `xml:"text"`
}
