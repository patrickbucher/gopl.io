package ex03_10

import "testing"

var tests = map[string]string{
	"0":                     "0",
	"+1":                    "+1",
	"-1":                    "-1",
	"+1.23":                 "+1.23",
	"-1.233":                "-1.233",
	"1.233456":              "1.233,456",
	"-1111.233456":          "-1,111.233,456",
	"-5431111.233456":       "-5,431,111.233,456",
	"2345234534.2345234545": "2,345,234,534.234,523,454,5",
}

var performanceTest = "5723984759232794823984793284"

func TestComma(t *testing.T) {
	for input, expected := range tests {
		got := Comma(input)
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
