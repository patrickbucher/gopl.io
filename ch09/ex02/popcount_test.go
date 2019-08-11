package popcount

import "testing"

var tests = map[uint64]int{
	0:   0, // 0
	1:   1, // 1
	2:   1, // 10
	3:   2, // 11
	4:   1, // 100
	5:   2, // 101
	6:   2, // 110
	7:   3, // 111
	8:   1, // 1000
	15:  4, // 1111
	16:  1, // 10000
	127: 7, // 1111111
	128: 1, // 10000000
}

func TestPopCount(t *testing.T) {
	for input, want := range tests {
		got := PopCount(input)
		if got != want {
			t.Errorf("PopCount(%d), expected %d, got %d", input, want, got)
		}
	}
}
