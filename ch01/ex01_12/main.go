package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // second color in palette
)

const (
	cycles  = 5
	res     = 0.001
	size    = 100
	nframes = 64
	delay   = 8
)

type params struct {
	Cycles  int     // number of complete x oscillator revolutions
	Res     float64 // angular resolution
	Size    int     // image canvas covers [-size..+size]
	Nframes int     // number of animation frames
	Delay   int     // delay between frames in 10ms units
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}
		p := params{Cycles: cycles, Res: res, Size: size, Nframes: nframes,
			Delay: delay}
		if c, ok := r.Form["cycles"]; ok {
			if ic, err := strconv.Atoi(c[0]); err == nil {
				p.Cycles = ic
			}
		}
		if r, ok := r.Form["res"]; ok {
			if rf, err := strconv.ParseFloat(r[0], 64); err == nil {
				p.Res = rf
			}
		}
		if s, ok := r.Form["size"]; ok {
			if si, err := strconv.Atoi(s[0]); err == nil {
				p.Size = si
			}
		}
		if n, ok := r.Form["nframes"]; ok {
			if ni, err := strconv.Atoi(n[0]); err == nil {
				p.Nframes = ni
			}
		}
		if d, ok := r.Form["delay"]; ok {
			if di, err := strconv.Atoi(d[0]); err == nil {
				p.Delay = di
			}
		}
		lissajous(w, p)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(w io.Writer, p params) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: p.Nframes}
	phase := 0.0 // phase difference
	for i := 0; i < p.Nframes; i++ {
		rect := image.Rect(0, 0, 2*p.Size+1, 2*p.Size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(p.Cycles*2)*math.Pi; t += p.Res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(p.Size+int(x*float64(p.Size)+0.5),
				p.Size+int(y*float64(p.Size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, p.Delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim) // NOTE: ignoring encoding errors
}
