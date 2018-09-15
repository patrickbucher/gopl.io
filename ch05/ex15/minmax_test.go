package ex15

import "testing"

var tests = []struct {
	Numbers []int64
	Min     int64
	Max     int64
	Found   bool
}{
	{[]int64{}, 0, 0, false},
	{[]int64{-1, 1}, -1, 1, true},
	{[]int64{1, 2, 3}, 1, 3, true},
	{[]int64{-2435, 237482734, 2345}, -2435, 237482734, true},
	{[]int64{-10000, -1000, -100, -10, 10, 100, 1000, 10000}, -10000, 10000, true},
}

func TestMinMax(t *testing.T) {
	for i := range tests {
		numbers := tests[i].Numbers
		expectedMin := tests[i].Min
		expectedMax := tests[i].Max
		expectedFound := tests[i].Found

		gotMin, gotFound := Min(numbers...)
		if gotMin != expectedMin || gotFound != expectedFound {
			t.Errorf("from %v expected min %d and found %v, was %d and %v\n",
				numbers, expectedMin, expectedFound, gotMin, gotFound)
		}

		gotMax, gotFound := Max(numbers...)
		if gotMax != expectedMax || gotFound != expectedFound {
			t.Errorf("from %v expected max %d and found %v, was %d and %v\n",
				numbers, expectedMax, expectedFound, gotMax, gotFound)
		}
	}
}
