package ex06

import (
	"testing"
)

const (
	space            = "\u0020"
	noBreakSpace     = "\u00a0"
	enQuad           = "\u2000"
	emQuad           = "\u2001"
	enSpace          = "\u2002"
	emSpace          = "\u2003"
	ideographicSpace = "\u3000"
)

var tests = []struct {
	Input    string
	Expected string
}{
	// two literal spaces squashed to one
	{"a  b", "a b"},
	// two regular spaces squashed to one
	{"a" + space + space + "b", "a b"},
	// leading/trailing spaces
	{space + noBreakSpace + "a" + enQuad + emQuad, " a "},
	// all combined
	{enSpace + emSpace + "a b  c   d" + ideographicSpace + space + noBreakSpace + "  e  ",
		" a b c d e "},
	// proper text
	{"This is  a   test.    Really.", "This is a test. Really."},
}

func TestSquash(t *testing.T) {
	for i := range tests {
		got := Squash(tests[i].Input)
		if got != tests[i].Expected {
			t.Errorf("expected %q, got %q\n", tests[i].Expected, got)
		}
	}
}
