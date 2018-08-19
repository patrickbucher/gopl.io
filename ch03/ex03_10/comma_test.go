package ex03_10

import "testing"

var tests = map[string]string{
	"0":          "0",
	"9":          "9",
	"12":         "12",
	"936":        "936",
	"5431":       "5,431",
	"83713":      "83,713",
	"937289":     "937,289",
	"1986345":    "1,986,345",
	"74958358":   "74,958,358",
	"937958123":  "937,958,123",
	"8789456631": "8,789,456,631",
}

var performanceTest = "5723984759232794823984793284"

func TestComma(t *testing.T) {
	test(t, Comma)
}

func TestRecursiveComma(t *testing.T) {
	test(t, RecursiveComma)
}

func test(t *testing.T, f func(string) string) {
	for input, expected := range tests {
		got := f(input)
		if got != expected {
			t.Errorf("expected: %q, got: %q\n", expected, got)
		}
	}
}

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Comma(performanceTest)
	}
}

func BenchmarkRecursiveComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RecursiveComma(performanceTest)
	}
}
