package eval

import (
	"testing"
)

func TestPrettyPrint(t *testing.T) {
	tests := []struct {
		input  string
		pretty string
		env    Env
	}{
		{"1+1*pow(2,3)+sqrt(16)", "1 + 1 * pow(2, 3) + sqrt(16)", Env{}},
		{"-x+y", "-x + y", Env{"x": 2, "y": 3}},
		// TODO: create additional test cases
	}

	for _, test := range tests {
		exprOrig, err := Parse(test.input)
		if err != nil {
			t.Error(err) // parse error
		}
		got := exprOrig.String()
		if got != test.pretty {
			t.Errorf("parsed '%s', expected pretty print '%s', got '%s'",
				test.input, test.pretty, got)
		}
		exprCopy, err := Parse(got)
		if err != nil {
			t.Errorf("pretty print '%s' unable to parse: %v", got, err)
		}
		orig := exprOrig.Eval(test.env)
		copy := exprCopy.Eval(test.env)
		if orig != copy {
			t.Errorf("original=%g, pretty-printed=%g", orig, copy)
		}
	}
}
