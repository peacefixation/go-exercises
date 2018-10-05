package main

import (
	"image/color"
	"image/png"
	"os"
)

func main() {
	palette := Palette{
		Name: "Red",
		Colors: map[int]color.Color{
			0: color.RGBA{100, 0, 0, 255},
			1: color.RGBA{175, 0, 0, 255},
			2: color.RGBA{255, 0, 0, 255},
		},
		NumColors: 3}

	img := Draw(400, 400, palette)

	// Save to out.png
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
