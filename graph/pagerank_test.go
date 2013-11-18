package graph

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func getPagerankTestGraph() *Graph {
	g := NewGraph()

	a := NewNode("A")
	b := NewNode("B")
	c := NewNode("C")

	g.Add(a)
	g.Add(b)
	g.Add(c)

	g.AddArc(a, b)
	g.AddArc(b, a)
	g.AddArc(b, c)
	g.AddArc(c, a)

	return g
}

// TestPageRank tests Graph.PageRank() using test data here:
// http://en.wikipedia.org/wiki/PageRank
func TestPageRank(t *testing.T) {
	g := getPagerankTestGraph()

	Convey("Before first PageRank, weights should be 1", t, func() {
		for node := range g.Nodes {
			So(node.Rank, ShouldAlmostEqual, 1.0)
		}
	})

	g.pageRankOnce()

	Convey("After first iteration", t, func() {
		So(g.Get("A").Rank, ShouldAlmostEqual, 1.425)
		So(g.Get("B").Rank, ShouldAlmostEqual, 1.0)
		So(g.Get("C").Rank, ShouldAlmostEqual, 0.575)
	})

	g.pageRankOnce()

	Convey("After second iteration", t, func() {
		So(g.Get("A").Rank, ShouldAlmostEqual, 1.06375)
		So(g.Get("B").Rank, ShouldAlmostEqual, 1.36125)
		So(g.Get("C").Rank, ShouldAlmostEqual, 0.575)
	})

	g.pageRankOnce()

	Convey("After third iteration", t, func() {
		So(g.Get("A").Rank, ShouldAlmostEqual, 1.217)
		So(g.Get("B").Rank, ShouldAlmostEqual, 1.054)
		So(g.Get("C").Rank, ShouldAlmostEqual, 0.728)
	})
}

func TestNormalizeRanks(t *testing.T) {
	g := getPagerankTestGraph()
	g.normalizeRanks()

	Convey("All weights should be set to 1/3", t, func() {
		for node := range g.Nodes {
			So(node.Rank, ShouldAlmostEqual, 1.0/3.0)
		}
	})
}
