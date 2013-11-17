package graph

const DEFAULT_DAMPING = 0.85
const DEFAULT_ITERATIONS = 15

func (graph *Graph) pageRankOnce() {
	for node, _ := range graph.Nodes {
		pointingToNode := graph.PointingTo(node)

		sumRanks := 0.0

		for _, neighbor := range pointingToNode {
			sumRanks += neighbor.Rank / float64(len(graph.Neighbors(neighbor)))
		}

		node.Rank = (1 - DEFAULT_DAMPING) + DEFAULT_DAMPING*sumRanks
	}
}

// PageRank runs a page rank algorithm on the graph n number of times.
// It uses DEFAULT_DAMPING, as the damping factor, which is a reasonable value.
//
// The algorithm is described here:
// http://en.wikipedia.org/wiki/PageRank
//
// and more simply here:
// http://stackoverflow.com/questions/3950627/python-implementation-of-pagerank
func (graph *Graph) PageRank(n int) {
	for i := 0; i < n; i++ {
		graph.pageRankOnce()
	}
}
