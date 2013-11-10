package graph

type Node struct {
	Name string
}

func NewNode(name string) *Node {
	return &Node{
		Name: name,
	}
}

func (self *Node) String() string {
	return self.Name
}
