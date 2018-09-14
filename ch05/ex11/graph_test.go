package ex11

import "testing"

var tests = []struct {
	Graph  map[string][]string
	Name   string
	Cyclic bool
}{
	{
		map[string][]string{
			"A": {"B", "D"},
			"B": {"C", "D"},
			"C": {"A"}, // loop back to A
			"D": {},
		},
		"Graph A",
		true,
	},
	{
		map[string][]string{
			"A": {"B", "D"},
			"B": {"C", "D"},
			"C": {"D"}, // dead end to D
			"D": {},
		},
		"Graph B",
		false,
	},
}

func TestIsCyclic(t *testing.T) {
	for i := range tests {
		name := tests[i].Name
		graph := tests[i].Graph
		expected := tests[i].Cyclic
		got := IsCyclic(graph)
		if got != expected {
			t.Errorf("graph '%s': expected to be cyclic: %v, was %v\n",
				name, expected, got)
		}
	}
}
