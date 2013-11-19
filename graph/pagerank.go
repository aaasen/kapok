package graph

const DEFAULT_DAMPING = 0.85
const DEFAULT_ITERATIONS = 15

// PageRank runs a page rank algorithm on the graph n number of times.
// It uses DEFAULT_DAMPING, as the damping factor, which is a reasonable value.
//
// The algorithm is described here:
// http://en.wikipedia.org/wiki/PageRank
//
// and more simply here:
// http://stackoverflow.com/questions/3950627/python-implementation-of-pagerank
func (graph *Graph) PageRank(n int) {
	graph.normalizeRanks()

	for i := 0; i < n; i++ {
		graph.pageRankOnce()
	}
}

// Sets all ranks to 1 / (number of nodes).
// This should only be used before running pagerank,
// when the weights mean nothing.
func (graph *Graph) normalizeRanks() {
	rank := 1.0

	if len(graph.Nodes) > 0 {
		rank = 1.0 / float64(len(graph.Nodes))
	}

	for node := range graph.Nodes {
		node.Rank = rank
	}
}

func (graph *Graph) pageRankOnce() {
	oldRanks := make(map[*Node]float64)

	for node := range graph.Nodes {
		oldRanks[node] = node.Rank
	}

	for node := range graph.Nodes {
		pointingToNode := graph.PointingTo(node)

		sumRanks := 0.0

		for _, neighbor := range pointingToNode {
			sumRanks += oldRanks[neighbor] / float64(len(graph.Neighbors(neighbor)))
		}

		node.Rank = (1 - DEFAULT_DAMPING) + DEFAULT_DAMPING*sumRanks
	}
}
