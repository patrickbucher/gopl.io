package ex07

import "testing"

var tests = []struct {
	Input    string
	Expected string
}{
	// empty set
	{"", ""},
	// palindrom
	{"sugus", "sugus"},
	// ASCII
	{"Das ist ein Test.", ".tseT nie tsi saD"},
	// multi-byte character
	{"Я", "Я"},
	// unicode text
	{"я ничего не знаю", "юанз ен огечин я"},
}

func TestReverse(t *testing.T) {
	for i := range tests {
		input := []byte(tests[i].Input)
		expected := []byte(tests[i].Expected)
		got := Reverse(input)
		if !equal(expected, got) {
			t.Errorf("expected %q, got %q\n", expected, got)
		}
	}
}

func equal(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
