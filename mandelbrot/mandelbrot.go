package main

import (
	"image"
	"image/color"
)

// algorithm from http://warp.povusers.org/Mandelbrot/

var colorMap map[int]color.Color
var numColors int

func init() {
	colorMap = make(map[int]color.Color)
	colorMap[0] = color.RGBA{100, 0, 0, 255}
	colorMap[1] = color.RGBA{175, 0, 0, 255}
	colorMap[2] = color.RGBA{255, 0, 0, 255}
	numColors = len(colorMap)
}

// Draw a Mandelbrot fractal with the given dimensions
func Draw(width, height int) *image.RGBA {
	// create an image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	ImageHeight := float64(height)
	ImageWidth := float64(width)

	MinRe := -2.0
	MaxRe := 1.0
	MinIm := -1.2
	MaxIm := MinIm + (MaxRe-MinRe)*ImageHeight/ImageWidth
	ReFactor := (MaxRe - MinRe) / (ImageWidth - 1)
	ImFactor := (MaxIm - MinIm) / (ImageHeight - 1)
	MaxIterations := 100

	for y := 0; y < height; y++ {
		cIm := MaxIm - float64(y)*ImFactor
		for x := 0; x < width; x++ {
			cRe := MinRe + float64(x)*ReFactor

			zRe := cRe
			zIm := cIm
			isInside := true
			n := 0
			for ; n < MaxIterations; n++ {
				zRe2 := zRe * zRe
				zIm2 := zIm * zIm

				if zRe2+zIm2 > 4 {
					isInside = false
					break
				}

				// z = z^2 + c
				zIm = 2*zRe*zIm + cIm
				zRe = zRe2 - zIm2 + cRe
			}

			if isInside {
				// draw a black pixel
				img.Set(x, y, color.RGBA{0, 0, 0, 255})
			} else {
				// draw a colored pixel
				img.Set(x, y, colorMap[n%numColors])
			}
		}
	}

	return img
}
