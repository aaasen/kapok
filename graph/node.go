package graph

// Node is an arbitrary element of the graph.
type Node struct {
	Name string
	Rank float64
}

// NewNode returns a node with the given name
func NewNode(name string) *Node {
	return &Node{
		Name: name,
		Rank: 0.0,
	}
}

// String returns a string representation of the node,
// which is currently just its name.
func (self *Node) String() string {
	return self.Name
}
