package main

import (
	"fmt"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	m "mathvisu/math"
	"os"
)

func main() {
	const nframes = 64
	var palette = []color.Color{color.White}
	for i := 1; i < nframes; i++ {
		palette = append(palette, color.RGBA{
			R: 255 - uint8(float64(i)/float64(nframes)*255),
			G: uint8(float64(i) / float64(nframes) * 255),
			B: 0,
			A: 255,
		})
	}

	f, err := os.Create("lissajous.gif")
	if err != nil {
		fmt.Println(err)
	}

	freq := rand.Float64() * 3.0
	phase := 0.0
	comp := func(t float64) (x, y float64) {
		x = math.Sin(t)
		y = math.Sin(freq*t + phase)
		return
	}
	post := func() {
		phase += 0.1
	}

	gifGen := m.GifMathematicalGenerator{
		XYCalculator:         comp,
		PostFrameCalculation: post,
		Nframes:              nframes,
		Delay:                8,
		Cycles:               5,
		Size:                 100,
		Res:                  0.0001,
		Palette:              palette,
		Thickness:            3,
	}

	if err := gif.EncodeAll(f, gifGen.Generate()); err != nil {
		panic(err)
	}
}
