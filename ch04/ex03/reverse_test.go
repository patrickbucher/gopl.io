package ex04_10

import "testing"

var tests = []struct {
	Input    [3]int
	Expected [3]int
}{
	{Input: [...]int{1, 2, 3}, Expected: [...]int{3, 2, 1}},
}

func TestReverse(t *testing.T) {
	for i := range tests {
		var input *[3]int
		input = &tests[i].Input
		if Reverse(input); !equal(input, &tests[i].Expected) {
			t.Errorf("expected %v, got %v\n", tests[i].Expected, tests[i].Input)
		}
	}
}

func equal(x, y *[3]int) bool {
	for i := 0; i < 3; i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
