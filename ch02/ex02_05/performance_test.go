package ex02_05

import (
	"testing"

	"gopl.io/ch02/popcount"
)

var tests = map[uint64]int{
	1:  1, // 1
	55: 5, // 110111
	99: 4, // 1100011
}

func TestOriginalPopCount(t *testing.T) {
	for k, v := range tests {
		if pc := popcount.PopCount(k); pc != v {
			t.Errorf("expected %d, got %d\n", v, pc)
		}
	}
}

func TestLoopedPopCount(t *testing.T) {
	for k, v := range tests {
		if pc := PopCount(k); pc != v {
			t.Errorf("expected %d, got %d\n", v, pc)
		}
	}
}

func BenchmarkOriginalPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range tests {
			popcount.PopCount(k)
		}
	}
}

func BenchmarkLoopedPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for k, _ := range tests {
			PopCount(k)
		}
	}
}
