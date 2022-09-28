package math

import (
	"image"
	"image/color"
	"image/gif"
	"math"
)

type GifMathematicalGenerator struct {
	XYCalculator                            func(t float64) (x, y float64)
	PostFrameCalculation                    func()
	Nframes, Delay, Cycles, Size, Thickness int
	Res                                     float64
	Palette                                 []color.Color
}

func (g *GifMathematicalGenerator) Generate() *gif.GIF {
	anim := gif.GIF{LoopCount: g.Nframes}
	for i := 1; i <= g.Nframes; i++ {
		rect := image.Rect(0, 0, int(2*g.Size+1), int(2*g.Size+1))
		img := image.NewPaletted(rect, g.Palette)
		for t := 0.0; t < float64(g.Cycles)*2*math.Pi; t += g.Res {
			x, y := g.XYCalculator(t)
			thickness(g.Thickness, img, float64(g.Size), x, y, len(g.Palette))
		}
		g.PostFrameCalculation()
		anim.Delay = append(anim.Delay, g.Delay)
		anim.Image = append(anim.Image, img)
	}
	return &anim
}

func thickness(thickness int, img *image.Paletted, size, x, y float64, lengthPalette int) {
	preCalcX := size + x*size + 0.5
	preCalcY := size + y*size + 0.5
	l := lengthPalette - 3
	for i := 0; i < thickness; i++ {
		scaledx := preCalcX + float64(i)
		coloridx := colorSelector(scaledx, size*2, l)
		img.SetColorIndex(int(scaledx), int(preCalcY), coloridx)
	}
}

func colorSelector(x, size float64, lengthPalette int) (idx uint8) {
	return uint8(x/size*float64(lengthPalette)) + 1
}
