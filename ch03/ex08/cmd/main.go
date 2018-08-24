// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/big"
	"os"

	"gopl.io/ch03/ex08"
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
				c = ex08.Mandelbrot64(z)
			case "complex128":
				z := complex(x, y)
				c = ex08.Mandelbrot128(z)
			case "bigFloat":
				z := complex(x, y)
				c = ex08.MandelbrotBigFloat(z)
			case "bigRat":
				xRat := big.NewRat(int64(px), int64(width*(xmax-xmin)+xmin))
				yRat := big.NewRat(int64(py), int64(width*(ymax-ymin)+ymin))
				c = ex08.MandelbrotBigRat(xRat, yRat)
			}
			img.Set(px, py, c)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: Ignoring errors
}
