package counter

import "testing"

var wordTests = []struct {
	input string
	words int
}{
	{"", 0},
	{"yes", 1},
	{"yes no", 2},
	{"what is it", 3},
	{"this is a test", 4},
	{"this is yet another test", 5},
}

func TestWordCounter(t *testing.T) {
	for i := range wordTests {
		var w WordCounter
		input, expected := wordTests[i].input, wordTests[i].words
		w.Write([]byte(input))
		if int(w) != expected {
			t.Errorf("word count for %q should be %v but was %v", input, expected, w)
		}
	}
}

var lineTests = []struct {
	input string
	lines int
}{
	{"", 0},
	{"this", 1},
	{"this\nis", 2},
	{"this\nis", 2},
	{"this\nis\na\ntest", 4},
}

func TestLineCounter(t *testing.T) {
	for i := range lineTests {
		var l LineCounter
		input, expected := lineTests[i].input, lineTests[i].lines
		l.Write([]byte(input))
		if int(l) != expected {
			t.Errorf("line count for %q should be %v but was %v", input, expected, l)
		}
	}
}
