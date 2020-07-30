package main

import (
	"fmt"
	"image/color"
	"log"

	"golang.org/x/image/font"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
)

// Tile represents a Sudoku tile.
//
// A Value of -1 represents a blank tile
type Tile struct {
	Value       int
	Diameter    int
	BorderWidth int
	Fill        color.Color
	Border      color.Color
	image       *ebiten.Image
}

const (
	borderWidth  = 2
	tileDiameter = width / boardDiameter
	// excess space left over from overlapping tiles
	// TODO: improve this math, as it's only working for widths/heights that are
	// multiples of 9
	excess   = tileDiameter*boardDiameter - (tileDiameter + (tileDiameter-borderWidth)*(boardDiameter-1))
	dpi      = 72
	fontSize = 30
	hinting  = font.HintingFull
)

var (
	tileFillColor     = color.White
	tileBorderColor   = color.Gray{100}
	activeBorderColor = color.RGBA{R: 54, G: 123, B: 235, A: 255}
	fontColor         = color.Black
	mplusNormalFont   font.Face
)

// https://ebiten.org/examples/font.html is a great example of how to load and
// use fonts
func init() {
	trueTypeFont, err := truetype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	mplusNormalFont = truetype.NewFace(trueTypeFont, &truetype.Options{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: hinting,
	})
}

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

// Image returns a tile's pictorial representation as an *ebiten.Image
//
// That tile has an outer border that is borderWidth pixels thick and colored
// border. The rest of the square is colored fill, and the whole square is diameter
// wide and tall.
func (t *Tile) Image() *ebiten.Image {
	if t.image == nil {
		t.Update()
	}
	return t.image
}

func (t *Tile) Update() {
	borderSquare := FilledRectangle(t.Diameter, t.Diameter, t.Border)
	innerDiameter := t.Diameter - t.BorderWidth*2
	innerSquare := FilledRectangle(innerDiameter, innerDiameter, t.Fill)
	translation := &ebiten.DrawImageOptions{}
	translation.GeoM.Translate(borderWidth, borderWidth)
	borderSquare.DrawImage(innerSquare, translation)
	// Only render tile text for values between 0 and 9
	if t.Value >= 0 && t.Value <= 9 {
		x, y := borderSquare.Size()
		number := fmt.Sprintf("%d", t.Value)
		fontDimensions := text.MeasureString(number, mplusNormalFont)
		// TODO: understand why centering the text with the +/- 2 constants appears
		// to work. I had trouble figuring out the font math here.
		text.Draw(borderSquare, number, mplusNormalFont, (x-fontDimensions.X)/2+2, (y+fontSize)/2-2, fontColor)
	}
	t.image = borderSquare
}

// FilledRectangle returns an *ebiten.Image that is filled with the given color
func FilledRectangle(width, height int, fill color.Color) *ebiten.Image {
	rectangle, _ := ebiten.NewImage(width, height, ebiten.FilterDefault)
	rectangle.Fill(fill)
	return rectangle
}
