package palindrome

import "testing"

var tests = []struct {
	sequence string
	expected bool
}{
	{"anna", true},
	{"sugus", true},
	{"maoam", true},
	{"abba", true},
	{"alba", false},
	{"manna", false},
	{"mamma", false},
	{"gaggi", false},
}

func TestIsPalindrome(t *testing.T) {
	for i := range tests {
		got := IsPalindrome(word(tests[i].sequence))
		if got != tests[i].expected {
			t.Errorf("IsPalindrome(%s)=%v, expected %v",
				tests[i].sequence, got, tests[i].expected)
		}
	}
}
