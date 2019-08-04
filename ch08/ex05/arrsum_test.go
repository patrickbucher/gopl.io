package arrsum

import "testing"

var (
	n     = int(1e8) * 2
	min   = 1
	max   = 999
	input = RandomNumbers(n, min, max)
)

func BenchmarkSerial(b *testing.B) {
	Sum(input)
}

func BenchmarkOneGoroutine(b *testing.B) {
	SumConcurrently(input, 1)
}

func BenchmarkTwoGoroutines(b *testing.B) {
	SumConcurrently(input, 2)
}

func BenchmarkThreeGoroutines(b *testing.B) {
	SumConcurrently(input, 3)
}

func BenchmarkFourGoroutines(b *testing.B) {
	SumConcurrently(input, 4)
}

func BenchmarkFiveGoroutines(b *testing.B) {
	SumConcurrently(input, 5)
}
