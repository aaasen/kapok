package graph

import (
	. "github.com/smartystreets/goconvey/convey"
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
	getTestGraph()
}

func TestPrint(t *testing.T) {
	t.Log(getTestGraph().String())
}

func TestAdjacent(t *testing.T) {
	g := getTestGraph()

	if !g.Adjacent(g.Get("A"), g.Get("B")) {
		t.Error("A should be adjacent to B, but isn't")
	}

	if g.Adjacent(g.Get("C"), g.Get("A")) {
		t.Error("C shouldn't be adjacent to A, but it is")
	}
}

func TestPointingTo(t *testing.T) {
	g := getTestGraph()

	Convey("A and B should point to C", t, func() {
		So(g.PointingTo(g.Get("C")), ShouldResemble, []*Node{g.Get("A"), g.Get("B")})
	})
}

func TestRemove(t *testing.T) {
	g := getTestGraph()
	g.Remove(g.Get("B"))

	if !g.Adjacent(g.Get("A"), g.Get("C")) &&
		g.Get("B") == nil {
		t.Fail()
	}
}

func TestAddArc(t *testing.T) {
	g := getTestGraph()
	g.AddArc(g.Get("B"), g.Get("A"))

	if !g.Adjacent(g.Get("B"), g.Get("A")) {
		t.Fail()
	}
}

func TestRemoveArc(t *testing.T) {
	g := getTestGraph()
	g.RemoveArc(g.Get("A"), g.Get("B"))

	if g.Adjacent(g.Get("A"), g.Get("B")) {
		t.Fail()
	}
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

func TestPagerank(t *testing.T) {
	g := getPagerankTestGraph()

	Convey("Before first PageRank, weights should be 1", t, func() {
		for node := range g.Nodes {
			So(node.Rank, ShouldEqual, 1)
		}
	})

	g.PageRank()

	Convey("After first iteration", t, func() {
		So(g.Get("A").Rank, ShouldEqual, 1.425)
		So(g.Get("B").Rank, ShouldEqual, 1)
		So(g.Get("C").Rank, ShouldEqual, 0.575)
	})

	g.PageRank()

	Convey("After second iteration", t, func() {
		So(g.Get("A").Rank, ShouldEqual, 1.06375)
		So(g.Get("B").Rank, ShouldEqual, 1.36125)
		So(g.Get("C").Rank, ShouldEqual, 0.575)
	})

	g.PageRank()

	Convey("After third iteration", t, func() {
		So(g.Get("A").Rank, ShouldEqual, 1.217)
		So(g.Get("B").Rank, ShouldEqual, 1.054)
		So(g.Get("C").Rank, ShouldEqual, 0.728)
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
