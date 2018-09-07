package main

import "image/color"

// Palette a color palette
type Palette struct {
	Name   string
	Colors map[int]color.Color
}
