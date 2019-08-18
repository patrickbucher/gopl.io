package pipeline

import "testing"

func TestPipeline(t *testing.T) {
	const message = "hello"
	ch := Pipeline(message, 100)
	if got := <-ch; got != message {
		t.Errorf("sent %s, got %s", message, got)
	}
}

func BenchmarkPipeline100(b *testing.B) {
	benchmarkPipeline(b, 100)
}

func BenchmarkPipeline1000(b *testing.B) {
	benchmarkPipeline(b, 1000)
}

func BenchmarkPipeline10000(b *testing.B) {
	benchmarkPipeline(b, 10000)
}

func BenchmarkPipeline100000(b *testing.B) {
	benchmarkPipeline(b, 100000)
}

func BenchmarkPipeline1000000(b *testing.B) {
	benchmarkPipeline(b, 1000000)
}

func BenchmarkPipeline2000000(b *testing.B) {
	benchmarkPipeline(b, 10000000)
}

func benchmarkPipeline(b *testing.B, n uint) {
	for i := 0; i < b.N; i++ {
		ch := Pipeline("hello", n)
		<-ch
	}
}
