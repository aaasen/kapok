// A memory-efficient graph for large datasets.
//
// Time complexities:
// Storage:		O(V + E)
// Add Node:	O(1)
// Add Arc:		O(1)
// Remove Node: O(E)
// Remove Arc: 	O(1)
package graph

import (
	"encoding/gob"
	"fmt"
	"io"
	"strings"
)

// Graph is memory-efficient labeled graph for large datasets.
type Graph struct {
	Nodes map[*Node][]*Node
	Names map[string]*Node
}

// NewGraph returns an empty graph.
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[*Node][]*Node),
		Names: make(map[string]*Node),
	}
}

// Get returns the node with the given name or null if none exists.
func (self *Graph) Get(name string) *Node {
	return self.Names[name]
}

// SafeGet returns the node with the given name
// or creates a new node with the given name if one does not exist.
func (self *Graph) SafeGet(name string) *Node {
	node := self.Names[name]

	if node == nil {
		node = NewNode(name)
		self.Add(node)
	}

	return node
}

// Add adds a node to the graph.
func (self *Graph) Add(node *Node) {
	self.Nodes[node] = make([]*Node, 0)
	self.Names[node.Name] = node
}

// Remove removes a node from the graph along with all arcs pointing to it.
func (self *Graph) Remove(node *Node) {
	delete(self.Nodes, node)
	delete(self.Names, node.Name)

	for neighbor, _ := range self.Nodes {
		self.RemoveArc(node, neighbor)
	}
}

// AddArc adds an arc between the origin and destination nodes.
func (self *Graph) AddArc(origin *Node, dest *Node) {
	self.Nodes[origin] = append(self.Nodes[origin], dest)
}

// RemoveArc removes the arc between the origin and destination nodes.
func (self *Graph) RemoveArc(origin *Node, toRemove *Node) {
	i := indexOf(self.Nodes[origin], toRemove)

	if i >= 0 {
		self.Nodes[origin] = append(self.Nodes[origin][:i], self.Nodes[origin][i+1:]...)
	}
}

// Adjacent returns true if the origin node has an arc to the destination node
// and false otherwise.
func (self *Graph) Adjacent(origin *Node, target *Node) bool {
	destIndex := indexOf(self.Nodes[origin], target)

	if destIndex == -1 {
		return false
	}

	return true
}

// String creates a string representation of the graph in the following form:
//
// node -> neighbor, neighbor, neighbor
func (self *Graph) String() string {
	str := ""

	for node, neighbors := range self.Nodes {
		str += fmt.Sprintf("%s -> (", node.String())

		neighborStrings := make([]string, len(neighbors))
		for i, neighbor := range neighbors {
			neighborStrings[i] = neighbor.String()
		}

		str += strings.Join(neighborStrings, ", ")
		str += ")\n"
	}

	return str
}

// Export writes a gob dump of the graph to the given writer
// and returns any error that it encounters.
func (self *Graph) Export(writer io.Writer) error {
	encoder := gob.NewEncoder(writer)

	return encoder.Encode(self)
}

// Import creates a graph from a gob dump in the given writer.
// and returns any errors that it encounters.
func Import(reader io.Reader) (error, *Graph) {
	decoder := gob.NewDecoder(reader)

	var graph Graph
	return decoder.Decode(&graph), &graph
}

// indexOf returns the index of the value in a slice or -1 if it is not found.
func indexOf(slice []*Node, target *Node) int {
	for i, node := range slice {
		if node == target {
			return i
		}
	}

	return -1
}
