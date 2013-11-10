package graph

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"strings"
)

type Graph struct {
	Nodes map[*Node]map[*Node]bool
	Names map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[*Node]map[*Node]bool),
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
	self.Nodes[node] = make(map[*Node]bool, 0)
	self.Names[node.Name] = node
}

func (self *Graph) Remove(node *Node) {
	delete(self.Nodes, node)
	delete(self.Names, node.Name)

	for _, neighbors := range self.Nodes {
		delete(neighbors, node)
	}
}

func (self *Graph) AddArc(origin *Node, dest *Node) {
	self.Nodes[origin][dest] = true
}

func (self *Graph) RemoveArc(origin *Node, toRemove *Node) {
	delete(self.Nodes[origin], toRemove)
}

func (self *Graph) Adjacent(x *Node, y *Node) bool {
	_, exists := self.Nodes[x][y]

	return exists
}

func (self *Graph) String() string {
	str := ""

	for node, neighbors := range self.Nodes {
		str += fmt.Sprintf("%s -> (", node.String())

		n := make([]string, 0)
		for neighbor, _ := range neighbors {
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
