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
			img.Set(px, py, mandelbrot(z))
		}
	}
	supersample(img)
	png.Encode(os.Stdout, img) // NOTE: Ignoring errors
}

func supersample(orig *image.RGBA) {
	const subpixels = 2 // 2 in both dimensions: 4 subpixels
	w, h := orig.Bounds().Dx(), orig.Bounds().Dy()
	wc, hc := w*subpixels, h*subpixels
	copy := image.NewRGBA(image.Rect(0, 0, wc, hc))
	for xc := 0; xc < wc; xc++ {
		for yc := 0; yc < hc; yc++ {
			copy.Set(xc, yc, orig.At(xc/subpixels, yc/subpixels))
		}
	}
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			var rs, gs, bs byte
			for xc := x * subpixels; xc < x*subpixels+subpixels; xc++ {
				for yc := y * subpixels; yc < y*subpixels+subpixels; yc++ {
					r, g, b, _ := copy.At(xc, yc).RGBA()
					rs += byte(r)
					gs += byte(g)
					bs += byte(b)
				}
			}
			rs /= subpixels * subpixels
			gs /= subpixels * subpixels
			bs /= subpixels * subpixels
			orig.Set(x, y, color.RGBA{R: rs, G: gs, B: bs, A: 0xff})
		}
	}
}

func mandelbrot(z complex128) color.Color {
	var v complex128
	for n := 0; n < len(palette.WebSafe); n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette.WebSafe[n]
		}
	}
	return color.Black
}
