package tree

import (
	"strconv"
	"strings"
)

type Tree struct {
	Value       int
	Left, Right *Tree
}

// Adds the value to the tree in sorted (ascending) order.
func Add(t *Tree, value int) *Tree {
	if t == nil {
		t = new(Tree)
		t.Value = value
		return t
	}
	if value < t.Value {
		t.Left = Add(t.Left, value)
	} else {
		t.Right = Add(t.Right, value)
	}
	return t
}

func (t *Tree) String() string {
	values := make([]string, 0)
	for _, v := range t.GetValues() {
		values = append(values, strconv.Itoa(v))
	}
	return "[" + strings.Join(values, ", ") + "]"
}

// Returns the values of the tree in sorted (ascending) order.
func (t *Tree) GetValues() []int {
	values := make([]int, 0)
	values = appendValues(values, t)
	return values
}

func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.Left)
		values = append(values, t.Value)
		values = appendValues(values, t.Right)
	}
	return values
}
