package ex09

import (
	"strings"
	"testing"
)

var tests = map[string]map[string]int{
	"Das ist ein Test":  map[string]int{"Das": 1, "ist": 1, "ein": 1, "Test": 1},
	"was ist das ist":   map[string]int{"was": 1, "ist": 2, "das": 1},
	"es ist wie es ist": map[string]int{"es": 2, "ist": 2, "wie": 1},
	"foo bar baz bum":   map[string]int{"foo": 1, "bar": 1, "baz": 1, "bum": 1},
}

func TestWordfreq(t *testing.T) {
	for input, expected := range tests {
		got := Wordfreq(strings.NewReader(input))
		if !equal(got, expected) {
			t.Errorf("input:\t%q\nexpected:\t%v\ngot:\t\t%v\n",
				input, expected, got)
		}
	}
}

func equal(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, va := range a {
		if vb, ok := b[k]; va != vb || !ok {
			return false
		}
	}
	return true
}
