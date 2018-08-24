package ex04_04

import "testing"

var tests = []struct {
	Input    []int
	Rotation int
	Expected []int
}{
	{[]int{}, 0, []int{}},
	{[]int{}, 1, []int{}},
	{[]int{1}, 0, []int{1}},
	{[]int{1}, 1, []int{1}},
	{[]int{1}, 2, []int{1}},
	{[]int{1, 2}, 0, []int{1, 2}},
	{[]int{1, 2}, 1, []int{2, 1}},
	{[]int{1, 2}, 2, []int{1, 2}},
	{[]int{1, 2, 3}, 0, []int{1, 2, 3}},
	{[]int{1, 2, 3}, 1, []int{2, 3, 1}},
	{[]int{1, 2, 3}, 2, []int{3, 1, 2}},
	{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
	{[]int{1, 2, 3}, 4, []int{2, 3, 1}},
	{[]int{1, 2, 3, 4}, 0, []int{1, 2, 3, 4}},
	{[]int{1, 2, 3, 4}, 1, []int{2, 3, 4, 1}},
	{[]int{1, 2, 3, 4}, 2, []int{3, 4, 1, 2}},
	{[]int{1, 2, 3, 4}, 3, []int{4, 1, 2, 3}},
	{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
	{[]int{1, 2, 3, 4}, 5, []int{2, 3, 4, 1}},
}

func TestRotate(t *testing.T) {
	for i := range tests {
		got := Rotate(tests[i].Input, tests[i].Rotation)
		if !equal(got, tests[i].Expected) {
			t.Errorf("rotate %v by %d: expected %v, got %v\n",
				tests[i].Input, tests[i].Rotation, tests[i].Expected, got)
		}
	}
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
