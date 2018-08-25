package ex05

import "testing"

var tests = []struct {
	Input    []string
	Expected []string
}{
	// negative tests (nothing changes)
	{[]string{}, []string{}},
	{[]string{"a"}, []string{"a"}},
	{[]string{"a", "b"}, []string{"a", "b"}},
	{[]string{"a", "b", "c"}, []string{"a", "b", "c"}},
	{[]string{"a", "b", "a"}, []string{"a", "b", "a"}},
	{[]string{"a", "b", "a", "b"}, []string{"a", "b", "a", "b"}},

	// positive tests (changes required)
	{[]string{"a", "a"}, []string{"a"}},
	{[]string{"a", "b", "b"}, []string{"a", "b"}},
	{[]string{"a", "b", "b", "b", "c"}, []string{"a", "b", "c"}},
	{[]string{"This", "is", "is", "a", "test."},
		[]string{"This", "is", "a", "test."}},
	{[]string{"This", " ", "is", " ", "a", " ", " ", "test."},
		[]string{"This", " ", "is", " ", "a", " ", "test."}},
}

func TestUniq(t *testing.T) {
	for i := range tests {
		if got := Uniq(tests[i].Input); !equal(tests[i].Expected, got) {
			t.Errorf("expected %v, got %v\n", tests[i].Expected, got)
		}
	}
}

func equal(a, b []string) bool {
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
