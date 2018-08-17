// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"
)

var t = flag.String("t", "complex64", "The data type to be used,"+
	"one of complex64, complex128, bigFloat or bigRat")

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	flag.Parse()
	if *t != "complex64" && *t != "complex128" && *t != "bigFloat" &&
		*t != "bigRat" {
		fmt.Fprintf(os.Stderr, "%q is not supported\n", *t)
		os.Exit(1)
	}
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// Image point (px, py) represents complex value z.
			var c color.Color
			switch *t {
			case "complex64":
				z := complex(float32(x), float32(y))
				c = mandelbrot64(z)
			case "complex128":
				z := complex(x, y)
				c = mandelbrot128(z)
			case "bigFloat":
				z := complex(x, y)
				c = mandelbrotBigFloat(z)
			}
			img.Set(px, py, c)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: Ignoring errors
}

func mandelbrot64(z complex64) color.Color {
	var v complex64
	for n := 0; n < len(palette.WebSafe); n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return palette.WebSafe[n]
		}
	}
	return color.Black
}

func mandelbrot128(z complex128) color.Color {
	var v complex128
	for n := 0; n < len(palette.WebSafe); n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette.WebSafe[n]
		}
	}
	return color.Black
}

func mandelbrotBigFloat(z complex128) color.Color {
	r, i := big.NewFloat(real(z)), big.NewFloat(imag(z))
	var vR, vI *big.Float
	var b *big.Float
	for n := 0; n < len(palette.WebSafe); n++ {
		// FIXME: segfault
		// v = v² + z
		vR = b.Mul(vR, vR).Add(vR, r)
		vI = b.Mul(vI, vI).Add(vI, i)
		// c² = a² + b²
		c := b.Sqrt(b.Add(b.Mul(vR, vR), b.Mul(vI, vI)))
		abs, _ := c.Float64()
		if abs > 2 {
			return palette.WebSafe[n]
		}
	}
	return color.Black
}
