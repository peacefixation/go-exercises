package main

import (
	"image/png"
	"os"
)

func main() {
	img := Draw(400, 400)

	// Save to out.png
	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
