package graph

import (
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

func BenchmarkAdd(b *testing.B) {
	g := NewGraph()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		g.Add(NewNode(string(i)))
	}
}

func BenchmarkAddArc(b *testing.B) {
	g := NewGraph()

	for i := 0; i < b.N; i++ {
		g.Add(NewNode(string(i)))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; i < b.N; i++ {
			g.AddArc(g.Get(string(i)), g.Get(string(j)))
		}
	}
}

func BenchmarkRemoveArc(b *testing.B) {
	g := NewGraph()

	for i := 0; i < b.N; i++ {
		g.Add(NewNode(string(i)))
	}

	for i := 0; i < b.N; i++ {
		for j := 0; i < b.N; i++ {
			g.AddArc(g.Get(string(i)), g.Get(string(j)))
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; i < b.N; i++ {
			g.RemoveArc(g.Get(string(i)), g.Get(string(j)))
		}
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
