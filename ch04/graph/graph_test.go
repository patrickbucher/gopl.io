package graph

import "testing"

func TestGraphFuncs(t *testing.T) {
	if hasEdge("A", "B") {
		t.Error("'A' and 'B' must not have an edge yet.")
	}
	addEdge("A", "B")
	if !hasEdge("A", "B") {
		t.Error("'A' and 'B' must have an edge now.")
	}
}
