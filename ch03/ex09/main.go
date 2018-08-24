package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"net/url"
	"strconv"
)

const (
	defaultWidth  = 1000
	defaultHeight = 1000
	defaultZoom   = 1.0
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "parse form: %v\n", err)
			return
		}
		width := getInt(r.Form, "w", defaultWidth)
		height := getInt(r.Form, "h", defaultHeight)
		zoom := getFloat(r.Form, "z", defaultZoom)
		mandelbrot(width, height, zoom, w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getInt(f url.Values, name string, fallback int) int {
	if str, ok := f[name]; ok {
		val, err := strconv.Atoi(str[0])
		if err == nil {
			return val
		}
	}
	return fallback
}

func getFloat(f url.Values, name string, fallback float64) float64 {
	if str, ok := f[name]; ok {
		val, err := strconv.ParseFloat(str[0], 64)
		if err == nil {
			return val
		}
	}
	return fallback
}

func mandelbrot(width, height int, zoom float64, w io.Writer) {
	const xmin, ymin, xmax, ymax = -2, -2, +2, +2
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	x0 := int((float64(width) * (zoom - 1)) / 2)
	y0 := int((float64(height) * (zoom - 1)) / 2)
	xE := int((float64(width) * (zoom + 1)) / 2)
	yE := int((float64(height) * (zoom + 1)) / 2)
	for px := x0; px < xE; px++ {
		for py := y0; py < yE; py++ {
			x := float64(px)/(float64(width)*zoom)*(xmax-xmin) + xmin
			y := float64(py)/(float64(height)*zoom)*(ymax-ymin) + ymin
			z := complex(x, y)
			c := mandelbrotColor(z)
			img.Set(px-x0, py-y0, c)
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

func mandelbrotColor(z complex128) color.Color {
	var v complex128
	for n := 0; n < len(palette.WebSafe); n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette.WebSafe[n]
		}
	}
	return color.Black
}
