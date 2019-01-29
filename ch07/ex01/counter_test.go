package counter

import "testing"

var tests = []struct {
	input string
	words int
	lines int
}{
	{"", 0, 0},
	{"this", 1, 1},
	{"this is", 2, 1},
	{"this is\na", 3, 2},
	{"this is\na test", 4, 2},
	{"this is\na stupid test", 5, 2},
	{"this is\nyet another\nstupid test", 6, 3},
}

func TestCounter(t *testing.T) {
	for i := range tests {
		var wc WordCounter
		var lc LineCounter
		input, words, lines := tests[i].input, tests[i].words, tests[i].lines
		wc.Write([]byte(input))
		if int(wc) != words {
			t.Errorf("word counter for %q should be %v but was %v", input, words, wc)
		}
		lc.Write([]byte(input))
		if int(lc) != lines {
			t.Errorf("line counter for %q should be %v but was %v", input, lines, lc)
		}
	}
}
