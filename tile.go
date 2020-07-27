package main

import (
	"image"
	"image/color"
	"image/draw"
)

// https://blog.golang.org/image-draw is a good read to understand what is
// happenening here.

// ColoredTile returns a square with a border.
//
// That square has an outer border that is borderWidth pixels thick and colored
// border. The rest of the square is colored fill, and the whole square is diameter
// wide and tall.
func ColoredTile(diameter, borderWidth int, fill, border color.Color) image.Image {
	borderSquare := FilledRectangle(diameter, diameter, border)
	innerDiameter := diameter - borderWidth*2
	innerSquare := FilledRectangle(innerDiameter, innerDiameter, fill)
	b := borderSquare.Bounds()
	innerTopLeft := b.Min.Add(image.Pt(borderWidth, borderWidth))
	ib := innerSquare.Bounds()
	// Convert the innerSquare's coordinate space to the borderSquare's
	r := image.Rectangle{innerTopLeft, innerTopLeft.Add(ib.Size())}
	draw.Draw(borderSquare, r, innerSquare, ib.Min, draw.Src)
	return borderSquare
}

// FilledRectangle returns a draw.Image that is filled with the given color
func FilledRectangle(width, height int, fill color.Color) draw.Image {
	outer := image.Rect(0, 0, width, height)
	rectangle := image.NewRGBA(outer)
	draw.Draw(rectangle, rectangle.Bounds(), image.NewUniform(fill), rectangle.Bounds().Min, draw.Src)
	return rectangle
}
