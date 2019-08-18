package arrsum

import "testing"

var (
	n     = int(1e6)
	min   = 1
	max   = 999
	input = RandomNumbers(n, min, max)
	want  = testSum(input)
)

func testSum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func TestSerial(t *testing.T) {
	got := Sum(input)
	if got != want {
		t.Errorf("TestSerial: expected %d, got %d", want, got)
	}
}

func TestConcurrently(t *testing.T) {
	got := SumConcurrently(input, 4)
	if got != want {
		t.Errorf("TestConcurrentlz: expected %d, got %d", want, got)
	}
}

func BenchmarkSerial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(input)
	}
}

func BenchmarkOneGoroutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumConcurrently(input, 1)
	}
}

func BenchmarkTwoGoroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumConcurrently(input, 2)
	}
}

func BenchmarkThreeGoroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumConcurrently(input, 3)
	}
}

func BenchmarkFourGoroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumConcurrently(input, 4)
	}
}

func BenchmarkFiveGoroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumConcurrently(input, 5)
	}
}

func BenchmarkTenGoroutines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumConcurrently(input, 10)
	}
}
