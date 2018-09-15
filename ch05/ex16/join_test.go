package ex16

import "testing"

var tests = []struct {
	Parts     []string
	Separator string
	Result    string
}{
	{
		[]string{"I", "don't", "know."},
		" ",
		"I don't know.",
	},
	{
		[]string{"un", "stopp", "able"},
		"",
		"unstoppable",
	},
	{
		[]string{
			"Roses are Red",
			"Violets are Blue",
			"Go is amazing",
			"And so are you"},
		"\n",
		"Roses are Red\nViolets are Blue\nGo is amazing\nAnd so are you",
	},
}

func TestJoin(t *testing.T) {
	for i := range tests {
		input := tests[i].Parts
		sep := tests[i].Separator
		expected := tests[i].Result
		got := Join(sep, input)
		if got != expected {
			t.Errorf("join %v with '%s', expected: '%s' but was '%s'\n",
				input, sep, expected, got)
		}
	}
}
