package main

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/hajimehoshi/ebiten"
)

// https://blog.golang.org/image-draw is a good read to understand what is
// happenening here.

// Tile represents a Sudoku tile
type Tile struct {
	Value       int
	Diameter    int
	BorderWidth int
	Fill        color.Color
	Border      color.Color
}

const (
	borderWidth  = 2
	tileDiameter = width / boardDiameter
	// excess space left over from overlapping tiles
	// TODO: improve this math, as it's only working for widths/heights that are
	// multiples of 9
	excess = tileDiameter*boardDiameter - (tileDiameter + (tileDiameter-borderWidth)*(boardDiameter-1))
)

var (
	tileFillColor   = color.White
	tileBorderColor = color.Gray{100}
)

// NewTile returns a Sudoku tile with default values
func NewTile(value int) Tile {
	return Tile{
		Value:       value,
		Diameter:    tileDiameter,
		BorderWidth: borderWidth,
		Fill:        tileFillColor,
		Border:      tileBorderColor,
	}
}

// Image returns a tile's pictorial representation as an ebiten.Image
//
// That square has an outer border that is borderWidth pixels thick and colored
// border. The rest of the square is colored fill, and the whole square is diameter
// wide and tall.
func (t Tile) Image() *ebiten.Image {
	borderSquare := FilledRectangle(t.Diameter, t.Diameter, t.Border)
	innerDiameter := t.Diameter - t.BorderWidth*2
	innerSquare := FilledRectangle(innerDiameter, innerDiameter, t.Fill)
	b := borderSquare.Bounds()
	innerTopLeft := b.Min.Add(image.Pt(t.BorderWidth, t.BorderWidth))
	ib := innerSquare.Bounds()
	// Convert the innerSquare's coordinate space to the borderSquare's
	r := image.Rectangle{innerTopLeft, innerTopLeft.Add(ib.Size())}
	draw.Draw(borderSquare, r, innerSquare, ib.Min, draw.Src)
	tileImage, _ := ebiten.NewImageFromImage(borderSquare, ebiten.FilterDefault)
	return tileImage
}

// FilledRectangle returns a draw.Image that is filled with the given color
func FilledRectangle(width, height int, fill color.Color) draw.Image {
	outer := image.Rect(0, 0, width, height)
	rectangle := image.NewRGBA(outer)
	draw.Draw(rectangle, rectangle.Bounds(), image.NewUniform(fill), rectangle.Bounds().Min, draw.Src)
	return rectangle
}
