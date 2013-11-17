package graph

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

func getTestGraph() *Graph {
	g := NewGraph()

	a := NewNode("A")
	b := NewNode("B")
	c := NewNode("C")

	g.Add(a)
	g.Add(b)
	g.Add(c)

	g.AddArc(a, b)
	g.AddArc(a, c)
	g.AddArc(b, c)
	g.AddArc(c, b)

	return g
}

func TestCreate(t *testing.T) {
	Convey("A graph should be created without a panic", t, func() {
		So(func() { getTestGraph() }, ShouldNotPanic)
	})
}

func TestPrint(t *testing.T) {
	t.Log(getTestGraph().String())
}

func TestAdjacent(t *testing.T) {
	g := getTestGraph()

	Convey("A should be adjacent to B and C should not be adjacent to A", t, func() {
		So(g.Adjacent(g.Get("A"), g.Get("B")), ShouldBeTrue)
		So(g.Adjacent(g.Get("C"), g.Get("A")), ShouldBeFalse)
	})
}

func TestNeighbors(t *testing.T) {
	g := getTestGraph()

	Convey("A should point to B and C", t, func() {
		So(g.Neighbors(g.Get("A")), ShouldResemble, []*Node{g.Get("B"), g.Get("C")})
	})
}

func TestPointingTo(t *testing.T) {
	g := getTestGraph()

	Convey("A and B should point to C", t, func() {
		So(g.PointingTo(g.Get("C")), ShouldResemble, []*Node{g.Get("A"), g.Get("B")})
	})
}

func TestRemove(t *testing.T) {
	g := getTestGraph()

	Convey("Removing B should work", t, func() {
		So(g.Get("B"), ShouldNotBeNil)

		So(func() { g.Remove(g.Get("B")) }, ShouldNotPanic)

		So(g.Get("B"), ShouldBeNil)
	})
}

func TestAddArc(t *testing.T) {
	g := getTestGraph()

	Convey("Adding an arc between B and A should make them adjacent", t, func() {
		So(func() { g.AddArc(g.Get("B"), g.Get("A")) }, ShouldNotPanic)
		So(g.Adjacent(g.Get("B"), g.Get("A")), ShouldBeTrue)
	})
}

func TestRemoveArc(t *testing.T) {
	g := getTestGraph()

	Convey("Removing an arc between A and B should make them non adjacent", t, func() {
		So(func() { g.RemoveArc(g.Get("A"), g.Get("B")) }, ShouldNotPanic)
		So(g.Adjacent(g.Get("A"), g.Get("B")), ShouldBeFalse)
	})
}

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

func shouldAlmostEqual(actual interface{}, expected ...interface{}) string {
	a := actual.(float64)
	b := actual.(float64)

	if math.Abs(a-b) < 0.000001 {
		return ""
	} else {
		return fmt.Sprintf("%v does not equal %v", a, b)
	}
}

// TestPageRank tests Graph.PageRank() using test data here:
// http://en.wikipedia.org/wiki/PageRank
func TestPageRank(t *testing.T) {
	g := getPagerankTestGraph()

	Convey("Before first PageRank, weights should be 1", t, func() {
		for node := range g.Nodes {
			So(node.Rank, shouldAlmostEqual, 1)
		}
	})

	g.pageRankOnce()

	Convey("After first iteration", t, func() {
		So(g.Get("A").Rank, shouldAlmostEqual, 1.425)
		So(g.Get("B").Rank, shouldAlmostEqual, 1)
		So(g.Get("C").Rank, shouldAlmostEqual, 0.575)
	})

	g.pageRankOnce()

	Convey("After second iteration", t, func() {
		So(g.Get("A").Rank, shouldAlmostEqual, 1.06375)
		So(g.Get("B").Rank, shouldAlmostEqual, 1.36125)
		So(g.Get("C").Rank, shouldAlmostEqual, 0.575)
	})

	g.pageRankOnce()

	Convey("After third iteration", t, func() {
		So(g.Get("A").Rank, shouldAlmostEqual, 1.217)
		So(g.Get("B").Rank, shouldAlmostEqual, 1.054)
		So(g.Get("C").Rank, shouldAlmostEqual, 0.728)
	})

}

func BenchmarkAdd(b *testing.B) {
	g := NewGraph()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		g.Add(NewNode(string(i)))
	}
}

func BenchmarkRemove(b *testing.B) {
	g := NewGraph()

	for i := 0; i < b.N; i++ {
		g.Add(NewNode(string(i)))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		g.Remove(g.Get(string(i)))
	}
}

func BenchmarkAdjacent(b *testing.B) {

}
