package ex08

import (
	"image/color"
	"image/color/palette"
	"math"
	"math/big"
	"math/cmplx"
)

func Mandelbrot64(z complex64) color.Color {
	var v complex64
	for n := 0; n < len(palette.WebSafe); n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return palette.WebSafe[n]
		}
	}
	return color.Black
}

func Mandelbrot128(z complex128) color.Color {
	var v complex128
	for n := 0; n < len(palette.WebSafe); n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette.WebSafe[n]
		}
	}
	return color.Black
}

// NOTE: This is way too slow to generate highres images!
func MandelbrotBigFloat(z complex128) color.Color {
	r, i := big.NewFloat(real(z)), big.NewFloat(imag(z))
	vR, vI := big.NewFloat(0.0), big.NewFloat(0.0)
	for n := 0; n < len(palette.WebSafe); n++ {
		// v = v² + z
		vR.Mul(vR, vR).Add(vR, r)
		vI.Mul(vI, vI).Add(vI, i)
		// c² = a² + b²
		f := big.NewFloat(0.0)
		c := f.Sqrt(f.Add(f.Mul(vR, vR), f.Mul(vI, vI)))
		abs, _ := c.Float64()
		if abs > 2 {
			return palette.WebSafe[n]
		}
	}
	return color.Black
}

// NOTE: This is way too slow to generate highres images!
func MandelbrotBigRat(x, y *big.Rat) color.Color {
	vX, vY := big.NewRat(0, 1), big.NewRat(0, 1)
	for n := 0; n < len(palette.WebSafe); n++ {
		vX.Mul(vX, vX).Add(vX, x)
		vY.Mul(vY, vY).Add(vY, y)
		a, _ := vX.Float64()
		b, _ := vY.Float64()
		c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
		if c > 2 {
			return palette.WebSafe[n]
		}
	}
	return color.Black
}
