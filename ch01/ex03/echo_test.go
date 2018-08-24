package echo

import (
	"strings"
	"testing"
)

var s string = "This is what not to do if a bird shits on you"

func BenchmarkEcho1(b *testing.B) {
	args := strings.Split(s, " ")
	for i := 0; i < b.N; i++ {
		Echo1(args)
	}
}
func BenchmarkEcho2(b *testing.B) {
	args := strings.Split(s, " ")
	for i := 0; i < b.N; i++ {
		Echo2(args)
	}
}
func BenchmarkEcho3(b *testing.B) {
	args := strings.Split(s, " ")
	for i := 0; i < b.N; i++ {
		Echo3(args)
	}
}
