package main

import "golang.org/x/tour/pic"

import (
	"image"
	"image/color"
)

type Image struct {
	minPoint, maxPoint image.Point
}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	var bound image.Rectangle
	bound.Min = i.minPoint
	bound.Max = i.maxPoint
	
	return bound
}

func (Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x + x), uint8(y + y), 255, 255}
}

func main() {
	m := Image{}
	m.minPoint = image.Point{10,200}
	m.maxPoint = image.Point{400,400}
	pic.ShowImage(m)
}