package graph

const DAMPING = 0.85

func (graph *Graph) PageRank() {
	for node, _ := range graph.Nodes {
		pointingToNode := graph.PointingTo(node)

		sumRanks := 0.0

		for _, neighbor := range pointingToNode {
			sumRanks += neighbor.Rank / float64(len(graph.Neighbors(neighbor)))
		}

		node.Rank = (1 - DAMPING) + DAMPING*sumRanks
	}
}
