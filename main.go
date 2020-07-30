package main

import (
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	width  = 630
	height = width
	// A Sudoku board is 9x9
	boardDiameter = 9
)

type Game struct {
	Board  [][]Tile
	active image.Point
}

func NewGame() *Game {
	g := &Game{}
	g.Board = make([][]Tile, boardDiameter)
	for i := 0; i < boardDiameter; i++ {
		g.Board[i] = make([]Tile, boardDiameter)
		for j := range g.Board[i] {
			g.Board[i][j] = NewTile(0)
		}
	}
	g.active = image.Pt(-1, -1)
	return g
}

func (g *Game) Tile(x, y int) *Tile {
	return &g.Board[y][x]
}

func (g *Game) DrawTile(x, y int, screen *ebiten.Image) {
	tile := g.Tile(x, y)
	slide := float64(tileDiameter - borderWidth)
	translation := &ebiten.DrawImageOptions{}
	xSlide, ySlide := 0.0, 0.0
	for i := 0; i < x; i++ {
		xSlide += slide
	}
	for i := 0; i < y; i++ {
		ySlide += slide
	}
	translation.GeoM.Translate(xSlide, ySlide)
	screen.DrawImage(tile.Image(), translation)
}

func (g *Game) ResetTile(x, y int) {
	tile := g.Tile(x, y)
	defaultTile := NewTile(tile.Value)
	*tile = defaultTile
}

func (g *Game) DrawBoard(screen *ebiten.Image) {
	for i := 0; i < boardDiameter; i++ {
		for j := 0; j < boardDiameter; j++ {
			g.DrawTile(i, j, screen)
		}
	}
	if g.active.X >= 0 {
		tile := g.Tile(g.active.X, g.active.Y)
		tile.Border = activeBorderColor
		g.DrawTile(g.active.X, g.active.Y, screen)
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	cursorX, cursorY := ebiten.CursorPosition()
	if In(cursorX, 0, width) && In(cursorY, 0, height) {
		if g.HasActiveTile() {
			g.ResetTile(g.active.X, g.active.Y)
		}
		x, y := cursorX/tileDiameter, cursorY/tileDiameter
		g.active = image.Pt(x, y)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawBoard(screen)
	fmt.Printf("FPS: %v; TPS: %v\n", ebiten.CurrentFPS(), ebiten.CurrentTPS())
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width - excess, height - excess
}

// HasActiveTile returns true when the board has an active tile selected.
//
// A negative Point indicates that there is no active tile
func (g *Game) HasActiveTile() bool {
	if g.active.X < 0 || g.active.Y < 0 {
		return false
	}
	return true
}

func In(x, min, max int) bool {
	if x >= min && x <= max {
		return true
	}
	return false
}

func main() {
	game := NewGame()
	ebiten.SetWindowSize(width-excess, height-excess)
	ebiten.SetWindowTitle("Sudoku")
	err := ebiten.RunGame(game)
	if err != nil {
		log.Fatal(err)
	}
}
