package tree

import (
	"strconv"
	"strings"
	"testing"
)

func TestEmptyString(t *testing.T) {
	var tree *Tree
	expected := "[]"
	got := tree.String()
	assertEquals(t, expected, got)
}

func TestShortString(t *testing.T) {
	var tree *Tree
	Add(tree, 2)
	Add(tree, 1)
	Add(tree, 3)
	expected := "[1, 2, 3]"
	got := tree.String()
	assertEquals(t, expected, got)
}

func TestLongString(t *testing.T) {
	const size = 100
	var tree *Tree
	values := make([]string, 0)
	for i := 0; i < 100; i++ {
		Add(tree, i)
		values = append(values, strconv.Itoa(i))
	}
	expected := "[" + strings.Join(values, ", ") + "]"
	got := tree.String()
	assertEquals(t, expected, got)
}

func assertEquals(t *testing.T, expected, got string) {
	if got != expected {
		t.Errorf("tree string expected to be %q but was %q", expected, got)
	}
}
