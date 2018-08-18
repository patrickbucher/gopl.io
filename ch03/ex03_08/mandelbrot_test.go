package ex03_08

import (
	"math/big"
	"testing"
)

func BenchmarkComplex64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mandelbrot64(complex64(1 + 1i))
	}
}

func BenchmarkComplex128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mandelbrot128(1 + 1i)
	}
}

func BenchmarkBigFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MandelbrotBigFloat(1 + 1i)
	}
}

func BenchmarkBigRat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := big.NewRat(int64(100), int64(3))
		y := big.NewRat(int64(100), int64(3))
		MandelbrotBigRat(x, y)
	}
}
