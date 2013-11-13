package graph

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"strings"
)

type Graph struct {
	Nodes map[*Node][]*Node
	Names map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[*Node][]*Node),
		Names: make(map[string]*Node),
	}
}

func (self *Graph) Get(name string) *Node {
	return self.Names[name]
}

func (self *Graph) SafeGet(name string) *Node {
	node := self.Names[name]

	if node == nil {
		node = NewNode(name)
		self.Add(node)
	}

	return node
}

func (self *Graph) Add(node *Node) {
	self.Nodes[node] = make([]*Node, 0)
	self.Names[node.Name] = node
}

func (self *Graph) Remove(node *Node) {
	delete(self.Nodes, node)
	delete(self.Names, node.Name)

	for neighbor, _ := range self.Nodes {
		self.RemoveArc(node, neighbor)
	}
}

func (self *Graph) AddArc(origin *Node, dest *Node) {
	self.Nodes[origin] = append(self.Nodes[origin], dest)
}

func (self *Graph) RemoveArc(origin *Node, toRemove *Node) {
	i := indexOf(self.Nodes[origin], toRemove)

	if i >= 0 {
		self.Nodes[origin] = append(self.Nodes[origin][:i], self.Nodes[origin][i+1:]...)
	}
}

func (self *Graph) Adjacent(origin *Node, target *Node) bool {
	destIndex := indexOf(self.Nodes[origin], target)

	if destIndex == -1 {
		return false
	}

	return true
}

func (self *Graph) String() string {
	str := ""

	for node, neighbors := range self.Nodes {
		str += fmt.Sprintf("%s -> (", node.String())

		n := make([]string, 0)
		for _, neighbor := range neighbors {
			n = append(n, neighbor.String())
		}

		str += strings.Join(n, ", ")
		str += ")\n"
	}

	return str
}

func (self *Graph) Export(writer io.Writer) {
	encoder := gob.NewEncoder(writer)

	err := encoder.Encode(self)

	if err != nil {
		log.Fatal("error exporting graph: ", err)
	}
}

func Import(reader io.Reader) *Graph {
	decoder := gob.NewDecoder(reader)

	var graph Graph
	err := decoder.Decode(&graph)

	if err != nil {
		log.Fatal("error importing graph: ", err)
	}

	return &graph
}

func indexOf(slice []*Node, target *Node) int {
	for i, node := range slice {
		if node == target {
			return i
		}
	}

	return -1
}
