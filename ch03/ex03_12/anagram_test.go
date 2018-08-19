package ex03_12

import "testing"

type testCase struct {
	A string
	B string
	E bool
}

var tests = []testCase{
	{A: "gaggi", B: "gagig", E: true},
	{A: "gaggi", B: "gakgi", E: false},
	{A: "foobar", B: "raboof", E: true},
	{A: "foobar", B: "foorar", E: false},
}

func TestIsAnagram(t *testing.T) {
	for i := range tests {
		a, b := tests[i].A, tests[i].B
		expected := tests[i].E
		got := IsAnagram(a, b)
		if got != expected {
			t.Errorf("anagram test for %q and %q should be %v but was %v\n",
				a, b, expected, got)
		}
	}
}
