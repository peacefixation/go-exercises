package main

import (
	"image"
	"image/color"
)

// algorithm from http://warp.povusers.org/Mandelbrot/

// Draw a Mandelbrot fractal with the given dimensions
func Draw(width, height int, palette Palette) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	imageHeight := float64(height)
	imageWidth := float64(width)

	minRe := -2.2
	maxRe := 0.8
	minIm := -1.5
	maxIm := minIm + (maxRe-minRe)*imageHeight/imageWidth
	reFactor := (maxRe - minRe) / (imageWidth - 1)
	imFactor := (maxIm - minIm) / (imageHeight - 1)
	maxIterations := 100000

	for y := 0; y < height; y++ {
		cIm := maxIm - float64(y)*imFactor
		for x := 0; x < width; x++ {
			cRe := minRe + float64(x)*reFactor

			inside, n := inside(x, y, cRe, cIm, maxIterations)

			if inside {
				img.Set(x, y, color.RGBA{0, 0, 0, 255})
			} else {
				img.Set(x, y, palette.Colors[n%palette.NumColors])
			}
		}
	}

	return img
}

func inside(x, y int, cRe, cIm float64, maxIterations int) (bool, int) {
	zRe := cRe
	zIm := cIm
	n := 0
	for ; n < maxIterations; n++ {
		zRe2 := zRe * zRe
		zIm2 := zIm * zIm

		if zRe2+zIm2 > 4 {
			return false, n
		}

		// z = z^2 + c
		zIm = 2*zRe*zIm + cIm
		zRe = zRe2 - zIm2 + cRe
	}

	return true, n
}
