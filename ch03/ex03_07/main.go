// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			f := func(z complex128) complex128 {
				return cmplx.Pow(z, 4) - 1
			}
			d := func(z complex128) complex128 {
				return 4 * cmplx.Pow(z, 3)
			}
			img.Set(px, py, newton(z, f, d))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: Ignoring errors
}

func newton(z complex128, f, d func(complex128) complex128) color.Color {
	const contrast = 15
	var x complex128
	for n := 0; n < len(palette.WebSafe); n++ {
		x = x - f(z)/d(z)
		if cmplx.Abs(x) > 2 {
			return palette.WebSafe[n]
		}
	}
	return color.Black
}
